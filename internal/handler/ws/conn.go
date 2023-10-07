package ws

import (
	"encoding/json"
	"fmt"
	"forum/internal/models"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type conn struct {
	clientID uint
	conn     *websocket.Conn
	mu       sync.Mutex
}

func (h *Handler) closeConn(c *conn) {
	c.mu.Lock()
	defer c.mu.Unlock()

	client, ok := h.clients[c.clientID]
	if !ok {
		return
	}

	for i := 0; i < len(client.conns); i++ {
		if client.conns[i] == c {
			client.conns[i].conn.Close()
			client.conns = append(client.conns[:i], client.conns[i+1:]...)
			break
		}
	}

	if len(client.conns) == 0 {
		delete(h.clients, client.ID)
	}
}

func (h *Handler) readPump(conn *conn) {
	defer h.closeConn(conn)
	conn.conn.SetReadLimit(maxMessageSize)
	conn.conn.SetReadDeadline(time.Now().Add(pongWait))

	for {
		event, err := conn.readEvent()
		if err != nil {
			log.Println(err)
			return
		}
		switch event.Type {
		case models.WSEventTypes.Message:
			err = h.newMessage(conn, &event)
		case models.WSEventTypes.MessagesRequest:
			err = h.getMessages(conn, &event)
		case models.WSEventTypes.ChatsRequest:
			// TODO
			fmt.Println("get chats request")
		case models.WSEventTypes.ReadMessageRequest:
			// TODO
			fmt.Println("read message request")
		case models.WSEventTypes.OnlineUsersRequest:
			// TODO
			fmt.Println("get online users")
		case models.WSEventTypes.TypingInRequest:
			// TODO
			fmt.Println("typing request")
		case models.WSEventTypes.PongMessage:
			err = conn.conn.SetReadDeadline(time.Now().Add(h.pongWait))
		}
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func (c *conn) readEvent() (models.WSEvent, error) {
	var event models.WSEvent
	_, msg, err := c.conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("error: %v", err)
		}
		return event, err
	}
	err = json.Unmarshal(msg, &event)
	return event, err
}

func (c *conn) writeJSON(data interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.conn.WriteJSON(data)
}
