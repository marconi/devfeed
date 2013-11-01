package realtime

import "github.com/marconi/devfeed/db"

// Project channel, holds project and roster to all subscribed members
type Channel struct {
    Project *db.Project
    Roster  *Roster
}

func NewChannel(proj *db.Project) *Channel {
    return &Channel{Project: proj, Roster: NewRoster()}
}
