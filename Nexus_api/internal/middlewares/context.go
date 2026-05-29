package middlewares

import (
	"errors"

	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
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

func GetUserRole(ctx *gin.Context) (schemas.Role, error) {
	userRole, exists := ctx.Get("user_role")
	if !exists {
		return "", errors.New("user_role not found.")
	}

	role, ok := userRole.(schemas.Role)
	if !ok {
		return "", errors.New("invalid user_role")
	}

	return role, nil
}
