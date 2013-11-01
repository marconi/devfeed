package db

import (
    "time"

    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
)

type Message struct {
    Id        bson.ObjectId `bson:"_id,omitempty"`
    AuthorId  *mgo.DBRef `bson:"author_id"`
    ProjectId *mgo.DBRef `json:"project_id"`
    Body      string `json:"body"`
    Created   time.Time `json:"created"`
}

func NewMessage(authorId, projId *mgo.DBRef, body string) *Message {
    return &Message{
        AuthorId: authorId,
        ProjectId: projId,
        Body: body,
        Created: time.Now().UTC(),
    }
}
