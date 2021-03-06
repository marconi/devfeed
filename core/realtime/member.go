package realtime

import (
	log "github.com/cihub/seelog"
	rds "github.com/garyburd/redigo/redis"
	"github.com/marconi/devfeed/core/realtime/websocket"
	"github.com/marconi/devfeed/libs/redis"
	"github.com/trevex/golem"
)

// Holds any form of membership connection for the user
type Member struct {
	ChannelName string
	Redis       *redis.RedisClient
	Websocket   *golem.Connection
}

func NewMember(conn *golem.Connection, channelName string) *Member {
	member := &Member{
		ChannelName: channelName,
		Redis:       redis.NewRedisClient(),
		Websocket:   conn,
	}

	// handle pubsub events
	go member.handlePubSub()

	return member
}

func (m *Member) handlePubSub() {
	m.Redis.Subscribe(m.ChannelName)
	defer m.Redis.Close()
	for {
		switch n := m.Redis.Receive().(type) {
		case rds.Message:
			payload := string(n.Data)
			log.Info("Message: ", n.Channel, " ", payload)
			m.Websocket.Emit("message", &websocket.Message{Body: payload})
		case rds.Subscription:
			log.Info("Subscription: ", n.Kind, " ", n.Channel)
			if n.Kind == "unsubscribe" {
				return
			}
		case error:
			log.Error(n)
			return
		}
	}
}
