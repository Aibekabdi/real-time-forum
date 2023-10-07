package ws

import (
	"context"
	"forum/internal/models"
)

type messageInput struct {
	ReceiverID uint   `json:"receiverID" validator:"required"`
	Content    string `json:"content,omitempty" validator:"required"`
}

func (h *Handler) newMessage(conn *conn, event *models.WSEvent) error {
	var (
		input    messageInput
		senderID uint
	)
	err := unmarshalEventBody(event, &input)
	if err != nil {
		return err
	}
	senderID = conn.clientID

	message, err := h.service.Chat.Create(context.TODO(), senderID, input.ReceiverID, input.Content)
	if err != nil {
		return err
	}

	h.sendEventToClient(&models.WSEvent{
		Type:       models.WSEventTypes.Message,
		Body:       message,
		ReceiverID: message.SenderID,
	})

	h.sendEventToClient(&models.WSEvent{
		Type:       models.WSEventTypes.Message,
		Body:       message,
		ReceiverID: message.ReceiverID,
	})

	return nil
}

type getMessagesInput struct {
	ReceiverID    uint `json:"receiverID" validator:"required"`
	LastMessageID uint `json:"lastMessageID"`
}

func (h *Handler) getMessages(conn *conn, event *models.WSEvent) error {
	var (
		input    getMessagesInput
		senderID uint
	)

	err := unmarshalEventBody(event, &input)
	if err != nil {
		return err
	}

	senderID = conn.clientID

	messages, err := h.service.Chat.GetMessages(context.TODO(), senderID, input.ReceiverID, input.LastMessageID)
	if err != nil {
		return err
	}

	h.sendEventToClient(&models.WSEvent{
		Type:       models.WSEventTypes.MessagesResponse,
		Body:       messages,
		ReceiverID: senderID,
	})
	
	return nil
}
