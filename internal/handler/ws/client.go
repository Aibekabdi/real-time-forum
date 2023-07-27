package ws

import (
	"forum/internal/models"
	"sync"
)

type client struct {
	conns []*conn
	mu    sync.Mutex
	models.User
}
