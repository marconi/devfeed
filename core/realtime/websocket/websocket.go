package websocket

import (
    "github.com/marconi/devfeed/db"
    log "github.com/cihub/seelog"
    "github.com/trevex/golem"
)

type WebSocketHandler struct{
    UserConn *UserConnMapping
}

func NewWebsocketHandler(ucm *UserConnMapping) *WebSocketHandler {
    return &WebSocketHandler{UserConn: ucm}
}

func (wsh *WebSocketHandler) Init(conn *golem.Connection, data *Init) {
    // load user
    user, err := db.GetUserById(data.UserId)
    if err != nil {
        log.Error("Unable to load user: ", err)
        return
    }

    log.Info("Init from: ", user.Id)
    wsh.UserConn.AddConn(user, conn)
}

func (wsh *WebSocketHandler) Close(conn *golem.Connection) {
    user, err := wsh.UserConn.GetUserByConn(conn)
    if err != nil {
        log.Error(err)
    } else {
        log.Info("Closing: ", user.Id)
        if err := wsh.UserConn.RemoveByUser(user); err != nil {
            log.Error("Unable to remove connection: ", err)
        }
    }
}
