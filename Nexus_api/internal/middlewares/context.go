package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserID(ctx *gin.Context) (uuid.UUID, error) {
	userId, exists := ctx.Get("user_id")
	if !exists {
		return uuid.Nil, errors.New("user_id not found.")
	}

	id, ok := userId.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("invalid user_id")
	}
	return id, nil
}
