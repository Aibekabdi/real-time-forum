package ws

import (
	"fmt"
	"forum/internal/models"
	"sync"
)

type client struct {
	conns []*conn
	mu    sync.Mutex
	models.User
}

func (h *Handler) sendEventToClient(event *models.WSEvent) {
	client, ok := h.clients[event.ReceiverID]
	if !ok {
		return
	}
	fmt.Println(event, client)
	client.mu.Lock()
	defer client.mu.Unlock()

	for i := 0; i < len(client.conns); i++ {
		conn := client.conns[i]

		err := conn.writeJSON(&event)
		if err != nil {
			h.closeConn(conn)
			continue
		}
	}
}
