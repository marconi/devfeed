package pivotal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var baseUrl = "https://www.pivotaltracker.com/services/v5"

type Error struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

type Project struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProjectMembership struct {
	Id        int `json:"id"`
	ProjectId int `json:"project_id"`
	PersonId  int `json:"person_id"`
}

type MembershipSummary struct {
	Id          int    `json:"id"`
	ProjectId   int    `json:"project_id"`
	ProjectName string `json:"project_name"`
}

// get the project from summary
func (ms *MembershipSummary) GetProject() *Project {
	return &Project{
		Id:   ms.ProjectId,
		Name: ms.ProjectName,
	}
}

type Timezone struct {
	OlsonName string `json:"olson_name"`
	Offset    string `json:"offset"`
}

type Me struct {
	Id       int                  `json:"id"`
	Name     string               `json:"name"`
	Initials string               `json:"initials"`
	Email    string               `json:"email"`
	ApiToken string               `json:"api_token"`
	Timezone *Timezone            `json:"time_zone"`
	Projects []*MembershipSummary `json:"projects"`
}

// get all projects based on membership summary's projects
func (me *Me) GetProjects() []*Project {
	var projects []*Project
	for _, memSummary := range me.Projects {
		projects = append(projects, memSummary.GetProject())
	}
	return projects
}

// get all memberships based on membership summary's projects
func (me *Me) GetMemberships() []*ProjectMembership {
	var memberships []*ProjectMembership
	for _, memSummary := range me.Projects {
		membership := &ProjectMembership{
			Id:        memSummary.Id,
			ProjectId: memSummary.ProjectId,
			PersonId:  me.Id,
		}
		memberships = append(memberships, membership)
	}
	return memberships
}

func Request(uri, method, token string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", baseUrl, uri), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-TrackerToken", token)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetMe(token string) (*Me, error) {
	res, err := Request("me", "GET", token)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		reqError := new(Error)
		json.Unmarshal(body, &reqError)
		return nil, errors.New(reqError.Error)
	}

	var me *Me
	json.Unmarshal(body, &me)
	return me, nil
}
