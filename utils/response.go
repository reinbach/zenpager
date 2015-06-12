package utils

type Message struct {
	Type    string
	Content string
}

type Response struct {
	Result   string
	Messages []Message
	ID       int64
}
