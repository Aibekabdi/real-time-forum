package http

import (
	"forum/internal/handler"
	"forum/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) userIdentify(c *gin.Context) {
	token := c.GetHeader(authorizationHeader)
	user, err := h.service.Auth.ParseToken(token)
	switch user.Role {
	case models.Roles.Guest:
		handler.ErrorResponse(c, http.StatusUnauthorized, "invalid type of role")
		return
	case models.Roles.User:
	default:
		handler.ErrorResponse(c, http.StatusUnauthorized, "invalid type of role")
		return
	}
	if err != nil {
		handler.ErrorResponse(c, http.StatusUnauthorized, "invalid type of role")
		return
	}
	if user.UserId == 0 {
		handler.ErrorResponse(c, http.StatusUnauthorized, "invalid user id")
		return
	}
	c.Set(models.UserCtx, user)
	c.Next()
}
