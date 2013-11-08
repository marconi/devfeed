package db

import (
	"time"

	"github.com/marconi/devfeed/core"
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

func (m *Message) GetAuthor() (*User, error) {
	var author *User

	// FindRef requires ObjectId so we have to reinitialize it here.
	// Probably too much, ugh! :(
	dbref := &mgo.DBRef{
		Collection: m.AuthorId.Collection,
		Id:         bson.ObjectIdHex(m.AuthorId.Id.(string)),
	}
	if err := core.Db.FindRef(dbref).One(&author); err != nil {
		return nil, err
	}
	return author, nil
}
