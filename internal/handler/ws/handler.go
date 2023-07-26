package ws

import (
	"forum/internal/handler"
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

func (h *Handler) WsHandler(c *gin.Context) {
	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		handler.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	log.Println(conn)
}
