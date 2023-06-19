package http

import (
	"forum/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPost(c *gin.Context) {
	var (
		input models.Post
		err   error
	)
	user, ok := c.Request.Context().Value(models.UserCtx).(models.UserToken)
	if !ok {
		h.errorResponse(c, http.StatusInternalServerError, "invalid user context")
		return
	}
	log.Println(user)

	if err = c.BindJSON(&input); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "invalid json body")
		return
	}
	//TODO add to db and validate

	c.JSON(http.StatusCreated, map[string]interface{}{
		"post_id": 0,
	})
}
