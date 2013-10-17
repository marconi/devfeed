package websocket

import (
	log "github.com/cihub/seelog"
	"github.com/trevex/golem"
)

var (
	WebSocket       = new(WebSocketHandler)
	UserConnMapping = make(map[*golem.Connection]string)
)

type WebSocketHandler struct{}

func (wsh *WebSocketHandler) Init(conn *golem.Connection, data *Init) {
	log.Info("Ping from: ", data.UserId)
	UserConnMapping[conn] = data.UserId
}

func (wsh *WebSocketHandler) Close(conn *golem.Connection) {
	if userId, ok := UserConnMapping[conn]; ok {
		log.Info("Closing: ", userId)
	}
}
