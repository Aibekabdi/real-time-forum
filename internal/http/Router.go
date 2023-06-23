package http

import (
	"forum/internal/service"

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

	api := router.Group("/api")
	{
		posts := api.Group("/posts")
		{
			posts.GET("/", h.getALLPosts)
			{
				posts.Use(h.userIdentify)
				posts.POST("/", h.createPost)
				posts.DELETE("/:id", h.deletePost)
			}
		}
	}
	return router
}
