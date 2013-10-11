package chat

type Subscribe struct {
	Channel string `json:"channel"`
}

type Message struct {
	Body string `json:"body"`
}
