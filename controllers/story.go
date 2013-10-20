package controllers

import (
	"net/http"
	"strconv"

	log "github.com/cihub/seelog"
	"github.com/marconi/devfeed/core"
	"github.com/marconi/devfeed/db"
	"github.com/marconi/devfeed/utils"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"labix.org/v2/mgo/bson"
)

type StoryController struct{}

func (c *StoryController) ReadMany(ctx context.Context) error {
	user, isLoggedIn := IsLoggedIn(ctx)
	if !isLoggedIn {
		return goweb.Respond.WithStatus(ctx, http.StatusUnauthorized)
	}

	projId, err := strconv.Atoi(ctx.QueryValue("project_id"))
	if err != nil {
		log.Error("Unable to get project id: ", err)
		return goweb.Respond.WithStatus(ctx, http.StatusBadRequest)
	}

	offset, err := strconv.Atoi(ctx.QueryValue("offset"))
	if err != nil {
		// if we didn't get a valid offset, start at 0
		offset = 0
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

	// query more stories
	stories, err := project.GetStories(utils.NewPaging(offset, core.Config.App.StoriesPaging))
	if err != nil {
		log.Error("Unable to get stories for project ", projId, " with offset ", offset, ": ", err)
	}

	return goweb.API.RespondWithData(ctx, stories)
}
