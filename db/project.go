package db

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/marconi/devfeed/core"
	"github.com/marconi/devfeed/libs/pivotal"
	"labix.org/v2/mgo/bson"
)

func SaveProjects(projects []*pivotal.Project) error {
	c := core.Db.C("projects")
	for _, proj := range projects {
		_, err := c.Upsert(bson.M{"id": proj.Id}, proj)
		if err != nil {
			return err
		}
	}
	return nil
}

func SaveMemberships(memberships []*pivotal.ProjectMembership) error {
	c := core.Db.C("memberships")
	for _, membership := range memberships {
		_, err := c.Upsert(bson.M{"id": membership.Id}, membership)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetAllProjects(token string) ([]*pivotal.Project, error) {
	res, err := pivotal.Request("projects", "GET", token)
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

	var projects []*pivotal.Project
	json.Unmarshal(body, &projects)
	return projects, nil
}

func GetProjectById(id int) (*pivotal.Project, error) {
	c := core.Db.C("projects")
	project := new(pivotal.Project)
	err := c.Find(bson.M{"id": id}).One(&project)
	if err != nil {
		return nil, err
	}
	return project, nil
}
