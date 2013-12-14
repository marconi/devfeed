package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	log "github.com/cihub/seelog"
	"github.com/eknkc/amber"
	"github.com/marconi/devfeed/controllers"
	"github.com/marconi/devfeed/core"
	"github.com/marconi/devfeed/core/realtime"
	"github.com/marconi/devfeed/db"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
)

func init() {
	core.LoadConfig()
	core.InitMongo()
	core.InitSessionStore(new(db.User))
	realtime.InitRealtime()
}

func index(ctx context.Context) error {
	r := ctx.HttpRequest()
	rw := ctx.HttpResponseWriter()
	session, _ := core.GetSession(r)

	// compile template
	t, err := amber.CompileFile(
		"templates/index.amber",
		amber.Options{core.Config.App.Debug, false},
	)
	if err != nil {
		panic(err)
	}

	user, isLoggedIn := controllers.IsLoggedIn(ctx)
	var projects []*db.Project
	if isLoggedIn {
		projects, err = user.GetProjects()
		if err != nil {
			log.Error("Unable to get projects: ", err)
		}
	}
	context := struct {
		Request  *http.Request
		Flashes  []interface{}
		User     *db.User
		Projects []*db.Project
		Debug    bool
	}{
		r,
		session.Flashes(),
		user,
		projects,
		core.Config.App.Debug,
	}
	if err = session.Save(r, rw); err != nil {
		log.Error("Unable to save session: ", err)
	}
	return t.Execute(rw, context)
}

func loggedin_only(ctx context.Context) error {
	if _, ok := controllers.IsLoggedIn(ctx); !ok {
		return goweb.Respond.WithPermanentRedirect(ctx, "/")
	}
	return index(ctx)
}

func anonymous_only(ctx context.Context) error {
	if _, ok := controllers.IsLoggedIn(ctx); ok {
		return goweb.Respond.WithPermanentRedirect(ctx, "/")
	}
	return index(ctx)
}

// validate login credential and store user in session
func login(ctx context.Context) error {
	r := ctx.HttpRequest()
	rw := ctx.HttpResponseWriter()
	session, _ := core.GetSession(r)

	email := ctx.PostValue("email")
	password := ctx.PostValue("password")
	user, err := db.Login(email, password)
	if err != nil {
		msg := struct {
			Body string `json:"body"`
			Type string `json:"type"`
		}{
			Body: err.Error(),
			Type: "alert",
		}
		data, err := json.Marshal(msg)
		if err != nil {
			log.Error("Unable to marshal: ", err)
			return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
		}
		rw.Header().Set("Content-Type", "application/json")
		return goweb.Respond.With(ctx, http.StatusUnauthorized, data)
	}

	session.Values["user"] = user
	if err = session.Save(r, rw); err != nil {
		log.Error("Unable to save session: ", err)
	}
	userInfo := struct {
		SessionID string `json:"sessionid"`
		Id        string `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		ApiToken  string `json:"apitoken"`
	}{
		SessionID: session.ID,
		Id:        user.Id.Hex(),
		Name:      user.Name,
		Email:     user.Email,
		ApiToken:  user.Person.ApiToken,
	}
	data, err := json.Marshal(userInfo)
	if err != nil {
		log.Error("Unable to marshal: ", err)
		return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json")
	return goweb.Respond.With(ctx, http.StatusOK, data)
}

// Validate register credential and save it
func register(ctx context.Context) error {
	rw := ctx.HttpResponseWriter()

	name := ctx.PostValue("name")
	email := ctx.PostValue("email")
	password := ctx.PostValue("password")

	fieldErrs, err := db.RegisterUser(name, email, password)
	if len(fieldErrs) > 0 {
		data, err := json.Marshal(fieldErrs)
		if err != nil {
			log.Error("Unable to marshal: ", err)
			return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
		}
		rw.Header().Set("Content-Type", "application/json")
		return goweb.Respond.With(ctx, http.StatusBadRequest, data)
	}

	if err != nil {
		log.Error(err)
		return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
	}

	// everything went fine
	msg := struct {
		Body string `json:"body"`
		Type string `json:"type"`
	}{
		Body: "Check your email to activate your account",
		Type: "success",
	}
	data, err := json.Marshal(msg)
	if err != nil {
		log.Error("Unable to marshal: ", err)
		return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json")
	return goweb.Respond.With(ctx, http.StatusOK, data)
}

// Logout the user by clearing user session
func logout(ctx context.Context) error {
	r := ctx.HttpRequest()
	session, _ := core.GetSession(r)
	_, ok := session.Values["user"]
	if ok {
		delete(session.Values, "user")
		if err := session.Save(r, ctx.HttpResponseWriter()); err != nil {
			log.Error("Unable to save session: ", err)
		}
	}
	return goweb.Respond.WithPermanentRedirect(ctx, "/")
}

// Checks if current request has a loggedin user
func isloggedin(ctx context.Context) error {
	user, ok := controllers.IsLoggedIn(ctx)
	if ok {
		session, _ := core.GetSession(ctx.HttpRequest())
		userInfo := struct {
			SessionID string `json:"sessionid"`
			Id        string `json:"id"`
			Name      string `json:"name"`
			Email     string `json:"email"`
			ApiToken  string `json:"apitoken"`
		}{
			SessionID: session.ID,
			Id:        user.Id.Hex(),
			Name:      user.Name,
			Email:     user.Email,
			ApiToken:  user.Person.ApiToken,
		}
		data, err := json.Marshal(userInfo)
		if err != nil {
			log.Error("Unable to marshal: ", err)
			return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
		}
		ctx.HttpResponseWriter().Header().Set("Content-Type", "application/json")
		return goweb.Respond.With(ctx, http.StatusOK, data)
	}
	return goweb.Respond.WithStatus(ctx, http.StatusUnauthorized)
}

// activate user account using activation key
func activate(ctx context.Context) error {
	r := ctx.HttpRequest()
	key := ctx.PathValue("key")
	session, _ := core.GetSession(r)

	if user := db.GetInactiveUserByKey(key); user != nil {
		if err := user.Activate(); err != nil {
			log.Error("Error activating user: ", err)
		}
		session.AddFlash("Activated! you can now login using your account")
	} else {
		session.AddFlash("Your activation key is used or has already expired")
	}
	if err := session.Save(r, ctx.HttpResponseWriter()); err != nil {
		log.Error("Unable to save session: ", err)
	}
	return goweb.Respond.WithPermanentRedirect(ctx, "/")
}

// resend activation code
func resend(ctx context.Context) error {
	email := ctx.PostValue("email")
	user, err := db.GetUserByEmail(email)

	// be passive and don't tell user about any errors
	if err == nil {
		if err = user.ResendActivationEmail(); err != nil {
			log.Error("Unable to re-send activation email: ", err)
		}
	}

	msg := struct {
		Body string `json:"body"`
		Type string `json:"type"`
	}{
		Body: "Check your email for a new validation link",
		Type: "success",
	}
	data, err := json.Marshal(msg)
	if err != nil {
		log.Error("Unable to marshal: ", err)
		return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
	}
	ctx.HttpResponseWriter().Header().Set("Content-Type", "application/json")
	return goweb.Respond.With(ctx, http.StatusOK, data)
}

// retrieve user password
func retrieve(ctx context.Context) error {
	email := ctx.PostValue("email")
	user, err := db.GetUserByEmail(email)

	// be passive and don't tell user about any errors
	if err == nil {
		if err = user.RetrievePassword(); err != nil {
			log.Error("Unable to retrieve password: ", err)
		}
	}

	msg := struct {
		Body string `json:"body"`
		Type string `json:"type"`
	}{
		Body: "Check your email for your new password",
		Type: "success",
	}
	data, err := json.Marshal(msg)
	if err != nil {
		log.Error("Unable to marshal: ", err)
		return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
	}
	ctx.HttpResponseWriter().Header().Set("Content-Type", "application/json")
	return goweb.Respond.With(ctx, http.StatusOK, data)
}

func settingsUpdate(ctx context.Context) error {
	r := ctx.HttpRequest()
	rw := ctx.HttpResponseWriter()
	session, _ := core.GetSession(r)

	d, err := ctx.RequestData()
	if err != nil {
		log.Error("Unable to retrieve request data: ", err)
		return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
	}

	log.Info(d)

	name := ctx.PostValue("name")
	email := ctx.PostValue("email")
	password := ctx.PostValue("password")
	apitoken := ctx.PostValue("apitoken")

	user, ok := session.Values["user"].(*db.User)
	if !ok {
		log.Error("No user found in session: ", session.ID)
		return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
	}

	fieldErrs, err := user.Update(name, email, password, apitoken)
	if len(fieldErrs) > 0 {
		data, err := json.Marshal(fieldErrs)
		if err != nil {
			log.Error("Unable to marshal: ", err)
			return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
		}
		rw.Header().Set("Content-Type", "application/json")
		return goweb.Respond.With(ctx, http.StatusBadRequest, data)
	}

	if err != nil {
		log.Error("Unable to update user settings: ", err)
		return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
	}

	// sync projects if there are no errors and an apitoken was provided
	if len(fieldErrs) == 0 && err == nil && apitoken != "" {
		wsConn, err := realtime.UserConn.GetConnByUser(user)
		if err != nil {
			log.Error("Unable to find websocket connection: ", err)
		} else {
			go user.SyncProjectsAndNotify(wsConn)
		}
	}

	// when no errors were found, update was successful
	// so we save the user to the session.
	session.Values["user"] = user
	if err = session.Save(r, rw); err != nil {
		log.Error("Unable to save session: ", err)
	}

	type UserExp struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		ApiToken string `json:"apitoken"`
	}
	type MessageExp struct {
		Body string `json:"body"`
		Type string `json:"type"`
	}
	msg := struct {
		User    UserExp    `json:"user"`
		Message MessageExp `json:"message"`
	}{
		UserExp{
			user.Name,
			user.Email,
			user.Person.ApiToken,
		}, MessageExp{
			"Settings has been updated",
			"success",
		},
	}
	data, err := json.Marshal(msg)
	if err != nil {
		log.Error("Unable to marshal: ", err)
		return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json")
	return goweb.Respond.With(ctx, http.StatusOK, data)
}

// run frontend tests
func tests(ctx context.Context) error {
	t, err := amber.CompileFile(
		"templates/tests.amber",
		amber.Options{true, false},
	)
	if err != nil {
		panic(err)
	}

	context := struct {
		Request *http.Request
	}{ctx.HttpRequest()}
	return t.Execute(ctx.HttpResponseWriter(), context)
}

func main() {
	// close mongo connection
	defer core.Mongo.Close()

	// close redis sessions
	defer core.Session.Close()

	// map static files path, only for debug
	if core.Config.App.Debug {
		goweb.MapStatic("/static", "static")
		goweb.Map("GET", "/tests", tests)
	}

	// views with functional handlers
	goweb.Map("GET", "/", index)
	goweb.Map("POST", "/login", login)
	goweb.Map("POST", "/register", register)
	goweb.Map("GET", "/logout", logout)
	goweb.Map("GET", "/isloggedin", isloggedin)
	goweb.Map("GET", "/activation/confirm/{key}", activate)
	goweb.Map("POST", "/activation/resend", resend)
	goweb.Map("POST", "/password/retrieve", retrieve)
	goweb.Map("POST", "/settings/update", settingsUpdate)

	// just to satisfy backbone pushstate
	goweb.Map("GET", "/login", anonymous_only)
	goweb.Map("GET", "/password/retrieve", anonymous_only)
	goweb.Map("GET", "/register", anonymous_only)
	goweb.Map("GET", "/activation/resend", anonymous_only)
	goweb.Map("GET", "/projects", loggedin_only)
	goweb.Map("GET", "/projects/{id}", loggedin_only)
	goweb.Map("GET", "/settings", loggedin_only)
	goweb.Map("GET", "/settings/pivotal", loggedin_only)

	// map controllers
	err := goweb.MapController("api/projects", new(controllers.ProjController))
	if err != nil {
		panic(err)
	}
	err = goweb.MapController("api/stories", new(controllers.StoryController))
	if err != nil {
		panic(err)
	}
	err = goweb.MapController("api/messages", new(controllers.MessageController))
	if err != nil {
		panic(err)
	}

	// run the server
	addr := fmt.Sprintf(":%s", strconv.Itoa(core.Config.App.Port))
	log.Info("Running at: ", addr)
	err = http.ListenAndServe(addr, goweb.DefaultHttpHandler())
	if err != nil {
		panic(err)
	}
}
