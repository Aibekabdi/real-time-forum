package models

import "time"

type Message struct {
	ID         uint      `json:"id"`
	SenderID   uint      `json:"senderid"`
	ReceiverID uint      `json:"receiverid"`
	Content    string    `json:"content"`
	SentAt     time.Time `json:"sentat"`
}

type WSMessage struct {
	SenderID   uint   `json:"senderid"`
	ReceiverID uint   `json:"receiverid"`
	Content    string `json:"content"`
}

type WSEvent struct {
	Type       string      `json:"type"`
	Body       interface{} `json:"body,omitempty"`
	ReceiverID uint        `json:"receiverid,omitempty"`
}

var WSEventTypes = struct {
	Message             string
	ChatsRequest        string
	ChatsResponse       string
	MessagesRequest     string
	MessagesResponse    string
	ReadMessageRequest  string
	ReadMessageResponse string
	OnlineUsersRequst   string
	OnlineUsersResponse string
	TypingInRequest     string
	TypingInResponse    string
	Error               string
	SuccessConnection   string
	PingMessage         string
	PongMessage         string
}{
	Message:             "message",
	ChatsRequest:        "chatsRequest",
	ChatsResponse:       "chatsResponse",
	MessagesRequest:     "messagesRequest",
	MessagesResponse:    "messagesResponse",
	ReadMessageRequest:  "readMessageRequest",
	ReadMessageResponse: "readMessageResponse",
	OnlineUsersRequst:   "onlineUsersRequest",
	OnlineUsersResponse: "onlineUsersResponse",
	TypingInRequest:     "typingInRequest",
	TypingInResponse:    "typingInResponse",
	Error:               "error",
	SuccessConnection:   "successConnection",
	PingMessage:         "pingMessage",
	PongMessage:         "pongMessage",
}
