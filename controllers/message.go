package controllers

import (
    "net/http"
    "errors"
    "fmt"
    "time"

    log "github.com/cihub/seelog"
    "labix.org/v2/mgo"
    "github.com/stretchr/goweb"
    "github.com/stretchr/goweb/context"
    "github.com/marconi/devfeed/db"
    "github.com/marconi/devfeed/core"
)

type MessageController struct{}

func (c *MessageController) ReadMany(ctx context.Context) error {
    user, isLoggedIn := IsLoggedIn(ctx)
    if !isLoggedIn {
        return goweb.Respond.WithStatus(ctx, http.StatusUnauthorized)
    }

    projId, err := strconv.Atoi()
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
    messages, err := project.GetRecentMessages(limit=core.Config.App.InitialChatMessages)
    if err != nil {
        log.Error("Unable to get recent messages for project ", projId, ": ", err)
    }

    return goweb.API.RespondWithData(ctx, messages)
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
        Id string `json:"id"`
        AuthorId string `json:"author_id"`
        ProjectId string `json:"project_id"`
        Body string `json:"body"`
        Created time.Time `json:"created"`
    }{
        message.Id.Hex(),
        message.AuthorId.Id.(string),
        message.ProjectId.Id.(string),
        message.Body,
        message.Created,
    }
    return goweb.API.RespondWithData(ctx, data)
}
