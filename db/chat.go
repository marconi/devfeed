package db

import (
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Message struct {
	Id        bson.ObjectId `bson:"_id,omitempty"json:"id"`
	AuthorId  *mgo.DBRef    `json:"author_id,string"`
	ProjectId *mgo.DBRef    `json:"project_id,string"` // TODO: find way to flatten DBRef
	Body      string        `json:"body"`
	Created   time.Time     `json:"created"`
}

func NewMessage(authorId, projId *mgo.DBRef, body string) *Message {
	return &Message{
		Id:        bson.NewObjectId(), // we generate id since insert doesn't fetch generated id
		AuthorId:  authorId,
		ProjectId: projId,
		Body:      body,
		Created:   time.Now().UTC(),
	}
}
