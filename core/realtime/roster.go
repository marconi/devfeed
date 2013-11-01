package realtime

import (
    "github.com/marconi/devfeed/db"
)

// Holds User:Websocket connection mapping
type Roster struct {
    Members map[*db.User]*Member
}

func NewRoster() *Roster {
    return &Roster{Members: make(map[*db.User]*Member)}
}

func (r *Roster) Add(user *db.User, member *Member) {
    r.Members[user] = member
}
