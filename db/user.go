package db

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"text/template"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
	log "github.com/cihub/seelog"
	"github.com/dchest/uniuri"
	"github.com/marconi/devfeed/core"
	"github.com/marconi/devfeed/libs/pivotal"
	"github.com/marconi/devfeed/libs/websocket"
	"github.com/marconi/devfeed/utils"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Id            bson.ObjectId `bson:"_id,omitempty"`
	Name          string
	Email         string
	Password      []byte
	IsActive      bool
	ActivationKey string
	Created       time.Time
	Person        *pivotal.Me
}

func NewInactiveUser(name, email, password string) (*User, error) {
	// generate password hash
	hp, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:          name,
		Email:         email,
		Password:      hp,
		IsActive:      false,
		ActivationKey: utils.GenerateKey(email),
		Created:       time.Now().UTC(),
		Person:        new(pivotal.Me),
	}, nil
}

func (u *User) GetId() string {
	userId, _ := u.Id.MarshalJSON()
	return string(userId)
}

func (u *User) SendActivationEmail() error {
	// we only send activation for inactive users
	if u.IsActive {
		return nil
	}

	var tplBuf bytes.Buffer
	var context = struct {
		User *User
		Link string
	}{
		User: u,
		Link: fmt.Sprintf("%s/activation/confirm/%s", core.Config.App.BaseUrl, u.ActivationKey),
	}
	t, err := template.ParseFiles("templates/emails/activation.txt")
	if err != nil {
		return err
	}
	err = t.Execute(&tplBuf, context)
	if err != nil {
		return err
	}

	return core.SendEmail(
		core.Config.Email.DefaultFrom,
		u.Email,
		"Activate your DevFeed account",
		tplBuf.String(),
	)
}

func (u *User) ResendActivationEmail() error {
	// only resend activation email for inactive users
	if u.IsActive {
		return nil
	}

	// generate new activation key and save it
	u.ActivationKey = utils.GenerateKey(u.Email)
	c := core.Db.C("users")
	if _, err := c.UpsertId(u.Id, &u); err != nil {
		return err
	}

	// resend the activation email
	go func() {
		if err := u.SendActivationEmail(); err != nil {
			log.Error("Unable to re-send activation email: ", err)
		}
	}()

	return nil
}

func (u *User) RetrievePassword() error {
	// only retrieve password of active users
	if !u.IsActive {
		return nil
	}

	// generate new password
	newPass := uniuri.New()
	hp, err := bcrypt.GenerateFromPassword([]byte(newPass), 0)
	if err != nil {
		return err
	}

	// save the hashed version of the password
	u.Password = hp
	c := core.Db.C("users")
	if _, err := c.UpsertId(u.Id, &u); err != nil {
		return err
	}

	// email the new password
	var tplBuf bytes.Buffer
	var context = struct {
		User     *User
		Password string
	}{
		User:     u,
		Password: newPass,
	}
	t, err := template.ParseFiles("templates/emails/password.txt")
	if err != nil {
		return err
	}
	err = t.Execute(&tplBuf, context)
	if err != nil {
		return err
	}

	go func() {
		if err := core.SendEmail(
			core.Config.Email.DefaultFrom,
			u.Email,
			"Your new DevFeed password",
			tplBuf.String(),
		); err != nil {
			log.Error("Unable to send new password email: ", err)
		}
	}()

	return nil
}

func (u *User) Activate() error {
	u.IsActive = true
	u.ActivationKey = ""

	c := core.Db.C("users")
	if _, err := c.UpsertId(u.Id, &u); err != nil {
		return err
	}
	return nil
}

func (u *User) Update(name, email, password, apitoken string) (map[string]string, error) {
	errs := make(map[string]string)
	if len(name) > 0 {
		u.Name = name
	}
	if len(email) > 0 {
		if !utils.IsValidEmail(email) {
			errs["email"] = "Must be a valid email"
		} else if user, exists := EmailExists(email); exists {
			if user.Id != u.Id {
				errs["email"] = "Email is already taken"
			}
		} else {
			u.Email = email
		}
	}
	if len(password) > 0 {
		if len(password) < 5 {
			errs["password"] = "Must be at least 5 characters"
		} else {
			// hash and store the hashed password
			hp, err := bcrypt.GenerateFromPassword([]byte(password), 0)
			if err != nil {
				return nil, err
			}
			u.Password = hp
		}
	}
	if len(apitoken) > 0 {
		// check if token is valid by fetching /me
		me, err := pivotal.GetMe(apitoken)
		if err != nil {
			log.Error("Unable to get me: ", err)
			errs["apitoken"] = "Invalid API Token"
		} else {
			u.Person = me
			u.Person.ApiToken = apitoken

			// sync projects
			go func() {
				if err := u.SyncProjects(); err != nil {
					log.Error("Error syncing projects: ", err)
				}
			}()
		}
	}

	// if there are field errors, return immediately
	if len(errs) > 0 {
		return errs, nil
	}

	c := core.Db.C("users")
	if _, err := c.UpsertId(u.Id, &u); err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *User) SyncProjects() error {
	// fetch projects
	projects, err := u.FetchProjects()
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to fetch projects: %s", err))
	}

	// build project memberships
	if err = u.BuildProjMemberships(); err != nil {
		return errors.New(fmt.Sprintf("Unable to build project memberships: %s", err))
	}

	// fetch stories of each project
	wsConn, err := websocket.UserConn.GetConnById(u.GetId())
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to websocket connection: %s", err))
	}

	c := core.Db.C("projects")
	for _, proj := range projects {
		// sync each project's stories concurrently, but shouldn't
		// spawn sub-goroutines inside since it'll exhaus PT's rate limit.
		go func() {
			// sync stories first
			if err := proj.SyncStories(u.Person.ApiToken); err != nil {
				log.Error("Unable to fetch stories of project ", proj.Id, " : ", err)
			} else {
				// mark project as synced
				proj.IsSynced = true
				if err := c.Update(bson.M{"id": proj.Id}, proj); err != nil {
					log.Error("Unable to update project ", proj.Id, " : ", err)
				}
				wsConn.Emit("project:synced", proj.Id)
			}
		}()
	}
	return nil
}

func (u *User) BuildProjMemberships() error {
	c := core.Db.C("memberships")
	for _, memsumm := range u.Person.Projects {
		projMem := &ProjectMembership{
			Id:        memsumm.Id,
			ProjectId: memsumm.ProjectId,
			PersonId:  u.Person.Id,
		}
		_, err := c.Upsert(bson.M{"id": projMem.Id}, projMem)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *User) FetchProjects() ([]*Project, error) {
	res, err := pivotal.Request("projects", "GET", u.Person.ApiToken)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		reqError := new(pivotal.Error)
		json.Unmarshal(body, &reqError)
		return nil, errors.New(reqError.Error)
	}

	var projects []*Project
	json.Unmarshal(body, &projects)

	// save projects
	c := core.Db.C("projects")
	for _, proj := range projects {
		_, err := c.Upsert(bson.M{"id": proj.Id}, proj)
		if err != nil {
			return nil, err
		}
	}
	return projects, nil
}

func (u *User) GetProjects() ([]*Project, error) {
	// get all project ids
	var projectIds []struct {
		ProjectId int `json:"projectid"`
	}
	mc := core.Db.C("memberships")
	err := mc.Find(bson.M{"personid": u.Person.Id}).Select(bson.M{"projectid": true}).All(&projectIds)
	if err != nil {
		return nil, err
	}

	if len(projectIds) > 0 {
		// flatten project ids
		var projIds []int
		for _, projId := range projectIds {
			projIds = append(projIds, projId.ProjectId)
		}

		// get all projects
		var projects []*Project
		pc := core.Db.C("projects")
		err = pc.Find(bson.M{"id": bson.M{"$in": projIds}}).All(&projects)
		if err != nil {
			return nil, err
		}
		return projects, nil
	}
	return nil, nil
}

// checks if user exists, is active and password matches the
// stored hashed version
func Login(email, password string) (*User, error) {
	genericErrMsg := "It's either your email or password is incorrect"
	inactiveErrMsg := "Your account is not yet active, check your email for activation link or <a href=\"/activation/resend\" class=\"resend-activation\">resend a new one</a>"
	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, errors.New(genericErrMsg)
	}
	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		return nil, errors.New(genericErrMsg)
	}
	if !user.IsActive {
		return nil, errors.New(inactiveErrMsg)
	}
	return user, nil
}

// checks if we already have a user with the same email in database
func EmailExists(email string) (*User, bool) {
	c := core.Db.C("users")
	user := new(User)
	err := c.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return nil, false
	}
	return user, true
}

func GetInactiveUserByKey(key string) *User {
	c := core.Db.C("users")
	user := new(User)
	err := c.Find(bson.M{"activationkey": key, "isactive": false}).One(&user)
	if err != nil {
		return nil
	}
	return user
}

func GetUserByEmail(email string) (*User, error) {
	c := core.Db.C("users")
	user := new(User)
	err := c.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// register new user and returns a map of errors and a
// generic error if it fails
func RegisterUser(name, email, password string) (map[string]string, error) {
	errs := make(map[string]string)

	// validate fields
	if len(name) == 0 {
		errs["name"] = "Name is required"
	}
	if len(email) == 0 {
		errs["email"] = "Email is required"
	} else if !utils.IsValidEmail(email) {
		errs["email"] = "Must be a valid email"
	} else if _, exists := EmailExists(email); exists {
		errs["email"] = "Email is already taken"
	}
	if len(password) == 0 {
		errs["password"] = "Password is required"
	} else if len(password) < 5 {
		errs["password"] = "Must be at least 5 characters"
	}

	// if there are field errors, return immediately
	if len(errs) > 0 {
		return errs, nil
	}

	// save the new user
	c := core.Db.C("users")
	user, err := NewInactiveUser(name, email, password)
	if err != nil {
		err = errors.New(fmt.Sprintf("Unable to create new inactive user: %s", err))
		return nil, err
	}

	if err = c.Insert(user); err != nil {
		err = errors.New(fmt.Sprintf("Unable to save user: %s", err))
		return nil, err
	}

	// send email activation
	go func() {
		if err = user.SendActivationEmail(); err != nil {
			log.Error("Unable to send activation email: %s", err)
		}
	}()

	return nil, nil
}
