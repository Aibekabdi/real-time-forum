package http

import (
	"errors"
	"forum/internal/models"

	"github.com/gin-gonic/gin"
)

func getUserFromCtx(c *gin.Context) (models.UserToken, error) {
	userInterface, ok := c.Get(models.UserCtx)
	if !ok {
		return models.UserToken{}, errors.New("invalid user context")
	}
	user, ok := userInterface.(models.UserToken)
	if !ok {
		return models.UserToken{}, errors.New("invalid type of user")
	}
	return user, nil
}
