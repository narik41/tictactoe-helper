package core

type MessageType string

const (
	MsgAuth MessageType = "AUTH"
)

type Message struct {
	MessageId       string      `json:"messageId"`
	Version         string      `json:"version"`
	MessageType     MessageType `json:"type"`
	IsAuthenticated bool        `json:"isAuthenticated"`
	Payload         interface{} `json:"payload"`
}

type TicTacToeMessage struct {
	MessageId string      `json:"messageId"`
	Version   string      `json:"version"`
	Timestamp int64       `json:"timestamp"`
	Payload   interface{} `json:"payload"`
}

type Version1MessagePayload struct {
	MessageType     string      `json:"messageType"`
	IsAuthenticated bool        `json:"isAuthenticated"`
	Payload         interface{} `json:"payload"`
}
