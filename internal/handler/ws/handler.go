package ws

import (
	"forum/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (h *Handler) wsHandler(c *gin.Context) {
	upgrader, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
