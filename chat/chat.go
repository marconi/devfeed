package chat

import (
	"log"

	"github.com/marconi/devfeed/libs/redis"
	rds "github.com/garyburd/redigo/redis"
	"github.com/trevex/golem"
)

// Member
type Member struct {
	ProjChan    string
	GolemConn   *golem.Connection
	RedisClient *redis.RedisClient
}

func NewMember(channel string, conn *golem.Connection) *Member {
	redisClient := redis.NewRedisClient()
	member := &Member{
		ProjChan:    channel,
		GolemConn:   conn,
		RedisClient: redisClient,
	}

	// handle pubsub events
	go member.handlePubSub()

	return member
}

func (m *Member) Destroy() {
	m.GolemConn.Close()
	m.RedisClient.Unsubscribe(m.ProjChan)
}

func (m *Member) handlePubSub() {
	m.RedisClient.Subscribe(m.ProjChan)
	defer m.RedisClient.Close()
	for {
		switch n := m.RedisClient.Receive().(type) {
		case rds.Message:
			payload := string(n.Data)
			log.Println("Message:", n.Channel, payload)
			m.GolemConn.Emit("message", &Message{Body: payload})
		case rds.Subscription:
			log.Println("Subscription:", n.Kind, n.Channel)
			if n.Kind == "unsubscribe" {
				return
			}
		case error:
			log.Println("error:", n)
			return
		}
	}
}

// Roster
type Roster struct {
	Members map[*golem.Connection]*Member
}

func NewRoster() *Roster {
	return &Roster{Members: make(map[*golem.Connection]*Member)}
}

func (r *Roster) Add(m *Member) int {
	r.Members[m.GolemConn] = m
	return r.Count()
}

func (r *Roster) Remove(conn *golem.Connection) int {
	if member, ok := r.Members[conn]; ok {
		member.Destroy()
		delete(r.Members, conn)
	}
	return r.Count()
}

func (r *Roster) Count() int {
	return len(r.Members)
}

// Chat handler
type ChatRoom struct {
	Name   string
	roster *Roster
}

func NewChatRoom(name string) *ChatRoom {
	return &ChatRoom{Name: name, roster: NewRoster()}
}

func (cr *ChatRoom) Join(m *Member) int {
	return cr.roster.Add(m)
}

func (cr *ChatRoom) Leave(conn *golem.Connection) int {
	return cr.roster.Remove(conn)
}
