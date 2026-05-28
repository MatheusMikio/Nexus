package handler

import (
	"net/http"

	"github.com/MatheusMikio/Nexus/internal/domain/dtos/auth"
	"github.com/gin-gonic/gin"
)

func SendSuccess(ctx *gin.Context, statusCode int, data any) {
	if statusCode == http.StatusNoContent {
		ctx.Status(statusCode)
		return
	}

	ctx.JSON(statusCode, gin.H{
		"data": data,
	})
}

func SendError(ctx *gin.Context, statusCode int, errors any) {
	ctx.JSON(statusCode, gin.H{
		"errors": errors,
	})
}

func SendAuthSuccess(ctx *gin.Context, statusCode int, accessToken string, expiresIn int64, user auth.AuthUser) {
	ctx.JSON(statusCode, gin.H{
		"accessToken": accessToken,
		"tokenType":   "Bearer",
		"expiresIn":   expiresIn,
		"user":        user,
	})
}
