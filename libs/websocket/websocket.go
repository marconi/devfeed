package websocket

import (
	"errors"
	"sync"

	log "github.com/cihub/seelog"
	"github.com/trevex/golem"
)

var (
	WebSocket = new(WebSocketHandler)
	UserConn  = NewUserConnMapping()
)

type UserConnMapping struct {
	UserConnId map[*golem.Connection]string
	UserIdConn map[string]*golem.Connection
	lock       sync.Mutex
}

func NewUserConnMapping() *UserConnMapping {
	return &UserConnMapping{
		UserConnId: make(map[*golem.Connection]string),
		UserIdConn: make(map[string]*golem.Connection),
	}
}

func (ucm *UserConnMapping) AddConn(id string, conn *golem.Connection) {
	ucm.lock.Lock()
	defer ucm.lock.Unlock()
	ucm.UserConnId[conn] = id
	ucm.UserIdConn[id] = conn
}

func (ucm *UserConnMapping) RemoveByConn(conn *golem.Connection) error {
	ucm.lock.Lock()
	defer ucm.lock.Unlock()
	id, ok := ucm.UserConnId[conn]
	if !ok {
		return errors.New("Connection not found.")
	}
	delete(ucm.UserIdConn, id)
	delete(ucm.UserConnId, conn)
	return nil
}

func (ucm *UserConnMapping) RemoveById(id string) error {
	ucm.lock.Lock()
	defer ucm.lock.Unlock()
	conn, ok := ucm.UserIdConn[id]
	if !ok {
		return errors.New("Connection not found.")
	}
	delete(ucm.UserConnId, conn)
	delete(ucm.UserIdConn, id)
	return nil
}

func (ucm *UserConnMapping) GetConnById(id string) (*golem.Connection, error) {
	conn, ok := ucm.UserIdConn[id]
	if !ok {
		return nil, errors.New("Connection not found.")
	}
	return conn, nil
}

func (ucm *UserConnMapping) GetIdByConn(conn *golem.Connection) (string, error) {
	id, ok := ucm.UserConnId[conn]
	if !ok {
		return "", errors.New("Connection not found.")
	}
	return id, nil
}

type WebSocketHandler struct{}

func (wsh *WebSocketHandler) Init(conn *golem.Connection, data *Init) {
	log.Info("Init from: ", data.UserId)
	UserConn.AddConn(data.UserId, conn)
}

func (wsh *WebSocketHandler) Close(conn *golem.Connection) {
	userId, err := UserConn.GetIdByConn(conn)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Closing: ", userId)
		if err := UserConn.RemoveById(userId); err != nil {
			log.Error("Unable to remove connection: ", err)
		}
	}
}
