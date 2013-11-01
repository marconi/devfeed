package websocket

import (
    "errors"
    "sync"

    "github.com/marconi/devfeed/db"
    "github.com/trevex/golem"
)

type UserConnMapping struct {
    ConnUserMap map[*golem.Connection]*db.User
    UserConnMap map[*db.User]*golem.Connection
    lock        sync.Mutex
}

func NewUserConnMapping() *UserConnMapping {
    return &UserConnMapping{
        ConnUserMap: make(map[*golem.Connection]*db.User),
        UserConnMap: make(map[*db.User]*golem.Connection),
    }
}

func (ucm *UserConnMapping) AddConn(user *db.User, conn *golem.Connection) {
    ucm.lock.Lock()
    defer ucm.lock.Unlock()
    ucm.ConnUserMap[conn] = user
    ucm.UserConnMap[user] = conn
}

func (ucm *UserConnMapping) RemoveByConn(conn *golem.Connection) error {
    ucm.lock.Lock()
    defer ucm.lock.Unlock()
    user, ok := ucm.ConnUserMap[conn]
    if !ok {
        return errors.New("Connection not found.")
    }
    delete(ucm.UserConnMap, user)
    delete(ucm.ConnUserMap, conn)
    return nil
}

func (ucm *UserConnMapping) RemoveByUser(user *db.User) error {
    ucm.lock.Lock()
    defer ucm.lock.Unlock()
    conn, ok := ucm.UserConnMap[user]
    if !ok {
        return errors.New("Connection not found.")
    }
    delete(ucm.ConnUserMap, conn)
    delete(ucm.UserConnMap, user)
    return nil
}

func (ucm *UserConnMapping) GetConnByUser(user *db.User) (*golem.Connection, error) {
    conn, ok := ucm.UserConnMap[user]
    if !ok {
        return nil, errors.New("Connection not found.")
    }
    return conn, nil
}

func (ucm *UserConnMapping) GetUserByConn(conn *golem.Connection) (*db.User, error) {
    user, ok := ucm.ConnUserMap[conn]
    if !ok {
        return nil, errors.New("Connection not found.")
    }
    return user, nil
}
