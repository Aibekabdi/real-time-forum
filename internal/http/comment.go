package http

import (
	"forum/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createComment(c *gin.Context) {
	var (
		input models.Comments
		err   error
	)

	user, err := getUserFromCtx(c)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if err = c.BindJSON(&input); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "invalid json body")
		return
	}
	input.Author.ID = user.UserId

	commentID, err := h.service.Comment.Create(c.Request.Context(), input)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"post_id": commentID,
	})
}

func (h *Handler) deleteComment(c *gin.Context) {
	user, err := getUserFromCtx(c)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	commentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	if err := h.service.Comment.Delete(c.Request.Context(), uint(commentID), user.UserId); err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "OK"})
}
