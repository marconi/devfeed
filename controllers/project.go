package controllers

import (
	"net/http"
	"strconv"

	log "github.com/cihub/seelog"
	"github.com/marconi/devfeed/core"
	"github.com/marconi/devfeed/db"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
)

func IsLoggedIn(ctx context.Context) (*db.User, bool) {
	session, err := core.GetSession(ctx.HttpRequest())
	user, ok := session.Values["user"]

	// if there's an error or we didn't get existing user,
	// then user is notauthorized.
	if err != nil || session.IsNew || !ok {
		return nil, false
	} else {
		switch t := user.(type) {
		case *db.User:
			return t, true
		default:
			return nil, false
		}
	}
}

type ProjController struct{}

func (c *ProjController) ReadMany(ctx context.Context) error {
	user, isLoggedIn := IsLoggedIn(ctx)
	if !isLoggedIn {
		return goweb.Respond.WithStatus(ctx, http.StatusUnauthorized)
	}
	projects, err := user.GetProjects()
	if err != nil {
		log.Error("Unable to get projects: ", err)
		return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
	}
	return goweb.API.RespondWithData(ctx, projects)
}

func (c *ProjController) Read(id string, ctx context.Context) error {
	_, isLoggedIn := IsLoggedIn(ctx)
	if !isLoggedIn {
		return goweb.Respond.WithStatus(ctx, http.StatusUnauthorized)
	}
	projectId, _ := strconv.Atoi(id)
	project, err := db.GetProjectById(projectId)
	if err != nil {
		log.Error("Unable to get project: ", err)
		return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
	}
	return goweb.API.RespondWithData(ctx, project)
}
