package chat

import (
	"log"
	// "net/http"
	"sync"

	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"github.com/trevex/golem"
)

var (
	ChatHandler = new(WebsocketHandler)
	ChatRooms   = NewRoomConnMapping()
)

// Room and Connection mapping
type RoomConnMapping struct {
	rooms       map[string]*ChatRoom
	connections map[*golem.Connection]*ChatRoom
	lock        sync.Mutex
}

func NewRoomConnMapping() *RoomConnMapping {
	return &RoomConnMapping{
		rooms:       make(map[string]*ChatRoom),
		connections: make(map[*golem.Connection]*ChatRoom),
	}
}

func (rcm *RoomConnMapping) Add(name string, conn *golem.Connection) int {
	rcm.lock.Lock()
	defer rcm.lock.Unlock()
	room, ok := rcm.rooms[name]
	if !ok {
		room = NewChatRoom(name)
		rcm.rooms[name] = room
	}
	rcm.connections[conn] = room
	return room.Join(NewMember(name, conn))
}

func (rcm *RoomConnMapping) Remove(conn *golem.Connection) (*ChatRoom, int) {
	rcm.lock.Lock()
	defer rcm.lock.Unlock()
	room := rcm.connections[conn]
	delete(rcm.connections, conn)
	return room, room.Leave(conn)
}

// Websocket handler
type WebsocketHandler struct{}

func InitChatHandler() {
	router := golem.NewRouter()
	router.On("subscribe", ChatHandler.Subscribe)
	router.OnClose(func(conn *golem.Connection) {
		room, count := ChatRooms.Remove(conn)
		log.Println(room.Name, "members:", count)
	})

	goweb.Map("/ws", func(ctx context.Context) error {
		handler := router.Handler()
		handler(ctx.HttpResponseWriter(), ctx.HttpRequest())
		return nil
	})
}

func (ch *WebsocketHandler) Subscribe(conn *golem.Connection, data *Subscribe) {
	count := ChatRooms.Add(data.Channel, conn)
	log.Println(data.Channel, "members:", count)
}
