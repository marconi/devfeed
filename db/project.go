package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/marconi/devfeed/core"
	"github.com/marconi/devfeed/libs/pivotal"
	"labix.org/v2/mgo/bson"
)

type Project struct {
	*pivotal.Project

	// devfeed specific fields
	IsSynced bool `json:"issynced"` // is project fully synced with pivotal
}

// fetch stories from pivotal
func (p *Project) FetchStories(token string) error {
	c := core.Db.C("stories")
	saveStories := func(stories []*pivotal.Story) error {
		for _, story := range stories {
			if _, err := c.Upsert(bson.M{"id": story.Id}, story); err != nil {
				return errors.New(fmt.Sprintf("Error saving stories: ", err))
			}
		}
		return nil
	}

	offset := 0
	limit := 100
	stories, total, err := p.GetStories(token, offset, limit)
	if err != nil {
		return errors.New(fmt.Sprintf("Error getting initial stories: ", err))
	}

	// save initial stories
	if err = saveStories(stories); err != nil {
		return err
	}

	// fetch more stories if there's still more
	offset += 1
	for total > limit*offset {
		from := limit * offset
		to := from + limit
		flimit := limit
		if total < to {
			excess := to - total
			to -= excess
			flimit = to - (offset * limit)
		}
		stories, total, err = p.GetStories(token, offset, flimit)
		if err != nil {
			return errors.New(fmt.Sprintf("Error getting stories: ", err))
		}
		if err = saveStories(stories); err != nil {
			return err
		}
		offset += 1
	}
	return nil
}

// fetch more info about the project from pivotal
func (p *Project) FetchMoreInfo(token string) error {
	url := fmt.Sprintf("projects/%d", p.Id)
	res, err := pivotal.Request(url, "GET", token)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		reqError := new(pivotal.Error)
		json.Unmarshal(body, &reqError)
		return errors.New(reqError.Error)
	}

	fmt.Println(string(body))

	// unserialize project info
	json.Unmarshal(body, &p)

	// save project
	c := core.Db.C("projects")
	_, err = c.Upsert(bson.M{"id": p.Id}, p)
	if err != nil {
		return err
	}
	return nil
}

type ProjectMembership struct {
	Id        int `json:"id"`
	ProjectId int `json:"project_id"`
	PersonId  int `json:"person_id"`
}

func SaveProjects(projects []*Project) error {
	c := core.Db.C("projects")
	for _, proj := range projects {
		_, err := c.Upsert(bson.M{"id": proj.Id}, proj)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetAllProjects(token string) ([]*Project, error) {
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

	var projects []*Project
	json.Unmarshal(body, &projects)
	return projects, nil
}

func GetProjectById(id int) (*Project, error) {
	c := core.Db.C("projects")
	project := new(Project)
	err := c.Find(bson.M{"project.id": id}).One(&project.Project)
	if err != nil {
		return nil, err
	}
	return project, nil
}
