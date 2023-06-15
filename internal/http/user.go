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
	if err := h.service.User.Create(c.Request.Context(), input); err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	//Todo validating user input and adding to db
	c.JSON(http.StatusOK, statusResponse{Status: "OK"})
}
