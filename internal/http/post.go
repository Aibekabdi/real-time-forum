package http

import (
	"forum/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPost(c *gin.Context) {
	var (
		input models.Post
		err   error
	)
	user, err := getUserFromCtx(c)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
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

func (h *Handler) deletePost(c *gin.Context) {
	user, err := getUserFromCtx(c)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	if err := h.service.Post.Delete(c.Request.Context(), uint(postID), user.UserId); err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "OK"})
}

func (h *Handler) getALLPosts(c *gin.Context) {
	var (
		posts []models.Post
		err   error
	)
	posts, err = h.service.Post.GetALL(c.Request.Context())
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, posts)
}
