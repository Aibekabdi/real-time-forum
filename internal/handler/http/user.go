package http

import (
	"forum/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserInfo(c *gin.Context) {
	userToken, err := getUserFromCtx(c)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userToken.Role == models.Roles.Guest {
		h.errorResponse(c, http.StatusUnauthorized, "invalid User Role")
		return
	}
	user, err := h.service.User.GetUserInfo(c.Request.Context(), userToken.UserId)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) updatePassword(c *gin.Context) {
	userToken, err := getUserFromCtx(c)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userToken.Role == models.Roles.Guest {
		h.errorResponse(c, http.StatusUnauthorized, "invalid User Role")
		return
	}
	var updatePsw models.UpdatePassword
	if err := c.BindJSON(&updatePsw); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "invalid json body")
		return
	}
	if err := h.service.User.UpdatePassword(c.Request.Context(), updatePsw, userToken.UserId); err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "OK"})
}
