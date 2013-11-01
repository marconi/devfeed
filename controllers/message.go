package controllers

import (
    "net/http"
    "errors"
    "fmt"

    log "github.com/cihub/seelog"
    "labix.org/v2/mgo"
    "github.com/stretchr/goweb"
    "github.com/stretchr/goweb/context"
    "github.com/marconi/devfeed/db"
    "github.com/marconi/devfeed/core"
)

type MessageController struct{}

func (c *MessageController) Create(ctx context.Context) error {
    user, isLoggedIn := IsLoggedIn(ctx)
    if !isLoggedIn {
        return goweb.Respond.WithStatus(ctx, http.StatusUnauthorized)
    }

    data, err := ctx.RequestData()
    if err != nil {
        return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
    }

    dataMap := data.(map[string]interface{})
    projId := dataMap["project_id"].(float64)
    body := dataMap["body"].(string)

    proj, err := db.GetProjectById(int(projId))
    if err != nil {
        log.Error("Unable to find project: ", err)
        return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)   
    }
    authorRef := &mgo.DBRef{Collection: "users", Id: user.Id.Hex()}
    projRef := &mgo.DBRef{Collection: "projects", Id: proj.ObjectId.Hex()}

    message := db.NewMessage(authorRef, projRef, body)
    mc := core.Db.C("messages")
    if err = mc.Insert(message); err != nil {
        log.Error(errors.New(fmt.Sprintf("Unable to save message: %s", err)))
        return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
    }
    return goweb.Respond.WithStatus(ctx, http.StatusCreated)
}
