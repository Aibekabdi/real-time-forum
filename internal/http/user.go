package http

import (
	"forum/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	// Validating user input and adding to db
	if err := h.service.User.Create(c.Request.Context(), input); err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "OK"})
}

func (h *Handler) signIn(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.service.User.SignIn(c.Request.Context(), input)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	}) // should be token
}
