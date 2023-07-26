package http

import (
	"forum/internal/handler"
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
		handler.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user.Role == models.Roles.Guest {
		handler.ErrorResponse(c, http.StatusUnauthorized, "invalid User Role")
		return
	}

	if err = c.BindJSON(&input); err != nil {
		handler.ErrorResponse(c, http.StatusBadRequest, "invalid json body")
		return
	}
	input.Author.ID = user.UserId

	postID, err := h.service.Post.Create(c.Request.Context(), input)
	if err != nil {
		handler.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"post_id": postID,
	})
}

func (h *Handler) deletePost(c *gin.Context) {
	user, err := getUserFromCtx(c)
	if err != nil {
		handler.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.ErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	if err := h.service.Post.Delete(c.Request.Context(), uint(postID), user.UserId); err != nil {
		handler.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, handler.StatusResponse{Status: "OK"})
}

func (h *Handler) getALLPosts(c *gin.Context) {
	var (
		posts []models.Post
		err   error
	)
	posts, err = h.service.Post.GetALL(c.Request.Context())
	if err != nil {
		handler.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *Handler) getPostByID(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.ErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	post, err := h.service.GetByID(c.Request.Context(), uint(postID))
	if err != nil {
		handler.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *Handler) likePostByID(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.ErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	user, err := getUserFromCtx(c)
	if err != nil {
		handler.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user.Role == models.Roles.Guest {
		handler.ErrorResponse(c, http.StatusUnauthorized, "invalid User Role")
		return
	}
	var input models.PostVote
	if err := c.BindJSON(&input); err != nil {
		handler.ErrorResponse(c, http.StatusBadRequest, "invalid json body")
		return
	}
	input.PostID = uint(postID)
	input.UserID = user.UserId

	if err := h.service.Post.InsertorDelete(c.Request.Context(), input); err != nil {
		handler.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, handler.StatusResponse{Status: "OK"})
}

func (h *Handler) getPostsByTag(c *gin.Context) {
	tagName := c.Param("tagName")
	posts, err := h.service.Post.GetALLByTag(c.Request.Context(), tagName)
	if err != nil {
		handler.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h *Handler) getPostsByUserID(c *gin.Context) {
	user, err := getUserFromCtx(c)
	if err != nil {
		handler.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if user.Role != models.Roles.User {
		handler.ErrorResponse(c, http.StatusUnauthorized, "invalid User Role")
		return
	}
	posts, err := h.service.GetALLByUserID(c.Request.Context(), user.UserId)
	if err != nil {
		handler.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, posts)
}
