package websocket

type Init struct {
	UserId string `json:"user_id"`
}

type ProjSub struct {
	ProjId string `json:"project_id"`
}

type Message struct {
	Body string `json:"body"`
}
