package core

import (
	"github.com/marconi/devfeed/libs/websocket"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"github.com/trevex/golem"
)

func InitRealtime() {
	router := golem.NewRouter()
	router.On("init", websocket.WebSocket.Init)
	router.OnClose(websocket.WebSocket.Close)

	goweb.Map("/ws", func(ctx context.Context) error {
		handler := router.Handler()
		handler(ctx.HttpResponseWriter(), ctx.HttpRequest())
		return nil
	})
}
