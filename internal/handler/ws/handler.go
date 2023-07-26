package ws

import (
	"forum/internal/service"
	"net/http"

	"github.com/gorilla/websocket"
)

type Handler struct {
	service  *service.Service
	upgrader websocket.Upgrader
}

func NewHandler(service *service.Service) *Handler {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	return &Handler{service: service, upgrader: upgrader}
}
