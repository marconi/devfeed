package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	log "github.com/cihub/seelog"
	"github.com/marconi/devfeed/core"
	"github.com/marconi/devfeed/db"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type MessageController struct{}

func (c *MessageController) ReadMany(ctx context.Context) error {
	user, isLoggedIn := IsLoggedIn(ctx)
	if !isLoggedIn {
		return goweb.Respond.WithStatus(ctx, http.StatusUnauthorized)
	}

	projId, err := strconv.Atoi(ctx.QueryValue("project_id"))
	if err != nil {
		log.Error("Unable to get project id: ", err)
		return goweb.Respond.WithStatus(ctx, http.StatusBadRequest)
	}

	// check if user actually is a member of project
	mc := core.Db.C("memberships")
	count, err := mc.Find(bson.M{"personid": user.Person.Id, "projectid": projId}).Count()
	if err != nil || count == 0 {
		log.Error("Unable to find membership of user ", user.Person.Id, " on project: ", projId)
		return goweb.Respond.WithStatus(ctx, http.StatusBadRequest)
	}

	// load project
	project, err := db.GetProjectById(projId)
	if err != nil {
		log.Error("Unable to get project: ", err)
		return goweb.Respond.WithStatus(ctx, http.StatusBadRequest)
	}

	// load recent messages
	messages, err := project.GetRecentMessages(core.Config.App.InitialChatMessages)
	if err != nil {
		log.Error("Unable to get recent messages for project ", projId, ": ", err)
		return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
	}

	// WARNING: hack to flatten DBRef for now
	type MsgFlat struct {
		Id        string    `json:"id"`
		AuthorId  string    `json:"author_id"`
		ProjectId string    `json:"project_id"`
		Body      string    `json:"body"`
		Created   time.Time `json:"created"`
	}
	var msgs []*MsgFlat
	for _, msg := range messages {
		msgFlat := &MsgFlat{
			Id:        msg.Id.Hex(),
			AuthorId:  msg.AuthorId.Id.(string),
			ProjectId: msg.ProjectId.Id.(string),
			Body:      msg.Body,
			Created:   msg.Created,
		}
		msgs = append(msgs, msgFlat)
	}

	return goweb.API.RespondWithData(ctx, msgs)
}

func (c *MessageController) Create(ctx context.Context) error {
	user, isLoggedIn := IsLoggedIn(ctx)
	if !isLoggedIn {
		return goweb.Respond.WithStatus(ctx, http.StatusUnauthorized)
	}

	data, err := ctx.RequestData()
	if err != nil {
		return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
	}

	// create message instance
	dataMap := data.(map[string]interface{})
	projId := dataMap["project_id"].(string)
	body := dataMap["body"].(string)
	authorRef := &mgo.DBRef{Collection: "users", Id: user.Id.Hex()}
	projRef := &mgo.DBRef{Collection: "projects", Id: projId}
	message := db.NewMessage(authorRef, projRef, body)

	// save the message
	mc := core.Db.C("messages")
	if err = mc.Insert(message); err != nil {
		log.Error(errors.New(fmt.Sprintf("Unable to save message: %s", err)))
		return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
	}

	data = struct {
		Id        string    `json:"id"`
		AuthorId  string    `json:"author_id"`
		ProjectId string    `json:"project_id"`
		Body      string    `json:"body"`
		Created   time.Time `json:"created"`
	}{
		message.Id.Hex(),
		message.AuthorId.Id.(string),
		message.ProjectId.Id.(string),
		message.Body,
		message.Created,
	}
	return goweb.API.RespondWithData(ctx, data)
}
