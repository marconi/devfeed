package controllers

import (
	"github.com/marconi/devfeed/core"
	"github.com/marconi/devfeed/db"
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
