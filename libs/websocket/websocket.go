package websocket

import (
	log "github.com/cihub/seelog"
	"github.com/trevex/golem"
)

var (
	WebSocket         = new(WebSocketHandler)
	UserConnIdMapping = make(map[*golem.Connection]string)
	UserIdConnMapping = make(map[string]*golem.Connection)
)

type WebSocketHandler struct{}

func (wsh *WebSocketHandler) Init(conn *golem.Connection, data *Init) {
	log.Info("Ping from: ", data.UserId)
	UserConnIdMapping[conn] = data.UserId
	UserIdConnMapping[data.UserId] = conn
}

func (wsh *WebSocketHandler) Close(conn *golem.Connection) {
	if userId, ok := UserConnIdMapping[conn]; ok {
		log.Info("Closing: ", userId)
	}
}
