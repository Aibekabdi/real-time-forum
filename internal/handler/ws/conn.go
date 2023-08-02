package ws

import (
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
		// _, msg, err := 
	}
}
