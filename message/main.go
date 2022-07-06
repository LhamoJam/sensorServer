package message

type Message struct {
	Timestamp int64  `json:"timestamp"`
	Content   []byte `json:"data"`
}

type MessageOperation interface {
	GetContent() []byte
}
