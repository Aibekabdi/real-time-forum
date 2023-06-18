package http

import (
	"forum/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	// Handlers is here

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentify)
	{
		// checking jwt tokens
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status": "OK",
			})
		}) // todo main forum place
	}
	return router
}
