package realtime

import (
    "strconv"
    "fmt"

    "github.com/marconi/devfeed/core/realtime/websocket"
    "github.com/marconi/devfeed/db"
    "github.com/stretchr/goweb"
    "github.com/stretchr/goweb/context"
    log "github.com/cihub/seelog"
    "github.com/trevex/golem"
)

var (
    UserConn  = websocket.NewUserConnMapping()
    WebSocket = websocket.NewWebsocketHandler(UserConn)
    Channels  = make(map[int]*Channel)
)

func ProjectSubscribe(conn *golem.Connection, data *websocket.ProjSub) {
    projId, _ := strconv.Atoi(data.ProjId)

    // create project channel if not yet present
    if _, ok := Channels[projId]; !ok {
        proj, err := db.GetProjectById(projId)
        if err != nil {
            log.Error("Unable to subscribe: ", err)
            return
        }
        Channels[projId] = NewChannel(proj)
    }

    // create new member and add to the roster of the
    // project where the user is subscribing.
    user, err := UserConn.GetUserByConn(conn)
    if err != nil {
        log.Error(err)
        return
    }
    channelName := fmt.Sprintf("project:%d", projId)
    Channels[projId].Roster.Add(user, NewMember(conn, channelName))
    log.Info("Subscribed to project: ", projId)
}

func InitRealtime() {
    router := golem.NewRouter()
    router.On("init", WebSocket.Init)
    router.On("project:subscribe", ProjectSubscribe)
    router.OnClose(WebSocket.Close)

    goweb.Map("/ws", func(ctx context.Context) error {
        handler := router.Handler()
        handler(ctx.HttpResponseWriter(), ctx.HttpRequest())
        return nil
    })
}
