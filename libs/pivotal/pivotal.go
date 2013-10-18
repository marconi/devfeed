package pivotal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var baseUrl = "https://www.pivotaltracker.com/services/v5"

type Error struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

type Timezone struct {
	OlsonName string `json:"olson_name"`
	Offset    string `json:"offset"`
}

type Project struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	TimeZone *Timezone `json:"time_zone"`

	// NOTE: unused pivotal fields
	// IterationLength              int        `json:"iteration_length"`
	// EnabledPlannedMode           bool       `json:"enabled_planned_mode"`
	// CurrentIterationNumber       int        `json:"current_iteration_number"`
	// CreatedAt                    *time.Time `json:"created_at"`
	// UpdatedAt                    *time.Time `json:"updated_at"`
	// PointScaleIsCustom           bool       `json:"point_scale_is_custom"`
	// Public                       bool       `json:"public"`
	// AccountId                    int        `json:"account_id"`
	// MembershipIds                []int      `json:"membership_ids"`
	// LabelIds                     []int      `json:"label_ids"`
	// EnableTasks                  bool       `json:"enable_tasks"`
	// Description                  string     `json:"description"`
	// InitialVelocity              int        `json:"initial_velocity"`
	// Version                      int        `json:"version"`
	// VelocityAveragedOver         int        `json:"velocity_averaged_over"`
	// NumberOfDoneIterationsToSHow int        `json:"number_of_done_iterations_to_show"`
	// EnableIncomingEmails         bool       `json:"enable_incoming_emails"`
	// AtomEnabled                  bool       `json:"atom_enabled"`
	// CurrentVelocity              int        `json:"current_velocity"`
	// PointScale                   string     `json:"point_scale"`
	// IntegrationIds               []int      `json:"integration_ids"`
	// BugsAndChoresAreEstimatable  bool       `json:"bugs_and_chores_are_estimatable"`
	// StartTime                    *time.Time `json:"start_time"`
	// EpicIds                      []int      `json:"epic_ids"`
	// ShowIterationsStartTime      *time.Time `json:"show_iterations_start_time"`
	// ProfileContent               string     `json:"profile_content"`
	// IterationOverrideNumbers     []int      `json:"iteration_override_numbers"`
	// WeekStartDay                 string     `json:"week_start_day"`
	// StartDate                    *time.Time `json:"start_date"`
	// HasGoogleDomain              bool       `json:"has_google_domain"`
	// StoryIds                     []int      `json:"story_ids"`
}

type membershipsummary struct {
	Id          int    `json:"id"`
	ProjectId   int    `json:"project_id"`
	ProjectName string `json:"project_name"`
}

type Me struct {
	Id       int                  `json:"id"`
	Name     string               `json:"name"`
	Initials string               `json:"initials"`
	Email    string               `json:"email"`
	ApiToken string               `json:"api_token"`
	TimeZone *Timezone            `json:"time_zone"`
	Projects []*membershipsummary `json:"projects"`
}

type Story struct {
	Id                     int        `json:"id"`
	ProjectId              int        `json:"project_id"`
	FollowerIds            []int      `json:"follower_ids"`
	UpdatedAt              *time.Time `json:"udpated_at"`
	CurrentState           string     `json:"current_state"` // accepted, delivered, finished, started, rejected, unstarted, unscheduled
	Name                   string     `json:"name"`
	CommentIds             []int      `json:"comment_ids"`
	Url                    string     `json:"url"`
	StoryType              string     `json:"story_type"` // feature, bug, chore, release
	LabelIds               []int      `json:"label_ids"`
	Description            string     `json:"description"`
	RequestedById          int        `json:"requested_by_id"`
	PlannedIterationNumber int        `json:"planned_iteration_number"`
	ExternalId             string     `json:"external_id"`
	Deadline               *time.Time `json:"deadline"`
	OwnedById              int        `json:"owned_by_id"`
	CreatedAt              *time.Time `json:"created_at"`
	Estimate               float64    `json:"estimate"`
	TaskIds                []int      `json:"task_ids"`
	IntegrationId          int        `json:"integration_id"`
	AcceptedAt             *time.Time `json:"accepted_at"`
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
