package http

import (
	"forum/internal/models"
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

	if user.Role == models.Roles.Guest {
		h.errorResponse(c, http.StatusUnauthorized, "invalid User")
		return
	}

	if err = c.BindJSON(&input); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "invalid json body")
		return
	}
	input.Author.Id = user.UserId
	//TODO add to db and validate
	postID, err := h.service.Post.Create(c.Request.Context(), input)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"post_id": postID,
	})
}
