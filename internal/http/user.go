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

	//Todo validating user input and adding to db
	c.JSON(http.StatusOK, input)
}
