package ws

import (
	"errors"
	"forum/internal/handler"
	"forum/internal/models"
	"forum/internal/service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	// time to read the next client's pong message
	pongWait = 60 * time.Second
	// time period to send pings to client
	pingPeriod = (pongWait * 9) / 10
	// time allowed to write a message to client
	writeWait = 10 * time.Second
	// max allowed connections for one user
	maxConnsForUser = 10
	// max message size allowed
	maxMessageSize = 512
	// I/O read buffer size
	readBufferSize = 1024
	// I/O write buffer size
	writeBufferSize = 1024
)

type Handler struct {
	clients    map[uint]*client
	service    *service.Service
	upgrader   websocket.Upgrader
	writeWait  time.Duration
	pongWait   time.Duration
	pingPeriod time.Duration
}

func NewHandler(service *service.Service) *Handler {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  readBufferSize,
		WriteBufferSize: writeBufferSize,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	return &Handler{
		clients:    make(map[uint]*client),
		service:    service,
		upgrader:   upgrader,
		writeWait:  writeWait,
		pongWait:   pongWait,
		pingPeriod: pingPeriod,
	}
}

func (h *Handler) WsHandler(c *gin.Context) {
	ws, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		handler.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	log.Println(ws.RemoteAddr(), ws.Subprotocol(), "conns")
	userToken, err := getUserFromCtx(c)
	if err != nil {
		handler.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	connection := &conn{conn: ws, clientID: userToken.UserId}
	curClient, ok := h.clients[connection.clientID]
	if !ok {
		user, err := h.service.User.GetUserInfo(c.Request.Context(), connection.clientID)
		if err != nil {
			connection.conn.WriteJSON(err)
			connection.conn.Close()
			return
		}
		curClient = &client{User: user}
		h.clients[connection.clientID] = curClient
	}
	if len(curClient.conns) == maxConnsForUser {
		conn := curClient.conns[0]
		conn.conn.WriteJSON("To many connections")
		h.closeConn(conn)
	}
	connection.conn.WriteJSON("Success Connection")
	// TODO: go routing conn read pump

	go h.readPump(connection)
	curClient.conns = append(curClient.conns, connection)
}

func getUserFromCtx(c *gin.Context) (models.UserToken, error) {
	userInterface, ok := c.Get(models.UserCtx)
	if !ok {
		return models.UserToken{}, errors.New("invalid user context")
	}
	user, ok := userInterface.(models.UserToken)
	if !ok {
		return models.UserToken{}, errors.New("invalid type of user")
	}
	return user, nil
}
