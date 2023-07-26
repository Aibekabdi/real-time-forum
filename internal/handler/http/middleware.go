package http

import (
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
		h.errorResponse(c, http.StatusUnauthorized, "invalid type of role")
		return
	case models.Roles.User:
	default:
		h.errorResponse(c, http.StatusUnauthorized, "invalid type of role")
		return
	}
	if err != nil {
		h.errorResponse(c, http.StatusUnauthorized, "invalid type of role")
		return
	}
	if user.UserId == 0 {
		h.errorResponse(c, http.StatusUnauthorized, "invalid user id")
		return
	}
	c.Set(models.UserCtx, user)
	c.Next()
}
