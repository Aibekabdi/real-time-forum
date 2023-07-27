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
