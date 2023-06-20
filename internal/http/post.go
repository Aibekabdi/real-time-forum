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
	userInterface, ok := c.Get(models.UserCtx)
	if !ok {
		h.errorResponse(c, http.StatusInternalServerError, "invalid user context")
		return
	}
	user, ok := userInterface.(models.UserToken)
	if !ok {
		h.errorResponse(c, http.StatusInternalServerError, "invalid type of user")
		return
	}

	if user.Role == models.Roles.Guest {
		h.errorResponse(c, http.StatusUnauthorized, "invalid User Role")
		return
	}

	if err = c.BindJSON(&input); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "invalid json body")
		return
	}
	input.Author.Id = user.UserId

	postID, err := h.service.Post.Create(c.Request.Context(), input)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"post_id": postID,
	})
}
