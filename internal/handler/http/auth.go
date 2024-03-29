package http

import (
	"forum/internal/handler"
	"forum/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		handler.ErrorResponse(c, http.StatusBadRequest, "invalid json body")
		return
	}

	// Validating user input and adding to db
	if err := h.service.Auth.Create(c.Request.Context(), input); err != nil {
		handler.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, handler.StatusResponse{Status: "Created"})
}

func (h *Handler) signIn(c *gin.Context) {
	var input models.SigningInput

	if err := c.BindJSON(&input); err != nil {
		handler.ErrorResponse(c, http.StatusBadRequest, "invalid json body")
		return
	}

	token, err := h.service.Auth.SignIn(c.Request.Context(), input)
	if err != nil {
		handler.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
