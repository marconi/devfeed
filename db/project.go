package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"

	log "github.com/cihub/seelog"
	"github.com/marconi/devfeed/core"
	"github.com/marconi/devfeed/libs/pivotal"
	"labix.org/v2/mgo/bson"
)

type Story struct {
	pivotal.Story `bson:",inline"`
}

type Task struct {
	pivotal.Task `bson:",inline"`
}

func (s *Story) FetchTasks(token string) ([]*Task, error) {
	url := fmt.Sprintf("projects/%d/stories/%d/tasks", s.ProjectId, s.Id)
	res, err := pivotal.Request(url, "GET", token)
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

	var tasks []*Task
	json.Unmarshal(body, &tasks)
	return tasks, nil
}

type Project struct {
	pivotal.Project `bson:",inline"`

	// devfeed specific fields
	IsSynced bool `json:"issynced"` // is project fully synced with pivotal
}

// return paginated stories, total stories or an error if there's any
func (p *Project) FetchStories(token string, offset, limit int) ([]*Story, int, error) {
	if offset > 0 {
		offset *= limit
	}
	url := fmt.Sprintf("projects/%d/stories?offset=%d&limit=%d", p.Id, offset, limit)
	res, err := pivotal.Request(url, "GET", token)
	if err != nil {
		return nil, 0, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}

	if res.StatusCode != 200 {
		reqError := new(pivotal.Error)
		json.Unmarshal(body, &reqError)
		return nil, 0, errors.New(reqError.Error)
	}

	total := 0
	totalHeader, ok := res.Header["X-Tracker-Pagination-Total"]
	if ok {
		total, err = strconv.Atoi(totalHeader[0])
		if err != nil {
			return nil, 0, err
		}
	}

	var stories []*Story
	json.Unmarshal(body, &stories)

	log.Debug("Fetching stories, project: ", p.Id, " offset: ", offset, " limit: ", limit, " total: ", len(stories))
	return stories, total, nil
}

// fetch stories from pivotal
func (p *Project) SyncStories(token string) error {
	sc := core.Db.C("stories")
	tc := core.Db.C("tasks")

	saveStories := func(stories []*Story) {
		for _, story := range stories {
			if _, err := sc.Upsert(bson.M{"id": story.Id}, story); err != nil {
				log.Error(errors.New(fmt.Sprintf("Error saving story: ", err)))
				continue
			}

			// we can't reliably use .TaskIds field since its
			// excluded on Story API response
			// if len(story.TaskIds) == 0 {
			// 	continue
			// }

			tasks, err := story.FetchTasks(token)
			if err != nil {
				log.Error(errors.New(fmt.Sprintf("Error fetching tasks: ", err)))
				continue
			}

			log.Debug("Found ", len(tasks), " tasks for story: ", story.Id)

			for _, task := range tasks {
				if _, err := tc.Upsert(bson.M{"id": task.Id}, task); err != nil {
					log.Error(errors.New(fmt.Sprintf("Error saving task: ", err)))
				}
			}
		}
	}

	offset := 0
	limit := 100
	stories, total, err := p.FetchStories(token, offset, limit)
	if err != nil {
		return errors.New(fmt.Sprintf("Error getting initial stories: ", err))
	}

	// save initial stories
	saveStories(stories)

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
		stories, total, err = p.FetchStories(token, offset, flimit)
		if err != nil {
			return errors.New(fmt.Sprintf("Error getting stories: ", err))
		}
		saveStories(stories)
		offset += 1
	}
	return nil
}

// fetch more info about the project from pivotal
func (p *Project) SyncMoreInfo(token string) error {
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

// get the stored  project stories
func (p *Project) GetStories() ([]*Story, error) {
	var stories []*Story
	c := core.Db.C("stories")
	if err := c.Find(bson.M{"projectid": p.Id}).All(&stories); err != nil {
		return nil, err
	}
	return stories, nil
}

type ProjectMembership struct {
	Id        int `json:"id"`
	ProjectId int `json:"project_id"`
	PersonId  int `json:"person_id"`
}

func GetProjectById(id int) (*Project, error) {
	c := core.Db.C("projects")
	project := new(Project)
	err := c.Find(bson.M{"id": id}).One(&project)
	if err != nil {
		return nil, err
	}
	return project, nil
}
