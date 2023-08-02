package ws

import (
	"forum/internal/models"
)

func (h *Handler) newMessage(conn *conn, event *models.WSEvent) error {
	var input models.WSMessage
	err := unmarshalEventBody(event, &input)
	if err != nil {
		return err
	}
	input.SenderID = conn.clientID
	//todo createMessage in database getting msg
	// h.sendEventToClient(event)
	return nil
}
