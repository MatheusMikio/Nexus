package middlewares

import (
	"net/http"
	"strings"

	"github.com/MatheusMikio/Nexus/internal/auth"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/handler"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if strings.TrimSpace(authHeader) == "" {
			handler.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "missing authorization header"))
			ctx.Abort()
			return
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			handler.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "invalid authorization header format"))
			ctx.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := auth.ValidateAccessToken(tokenString)
		if err != nil {
			handler.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "invalid or expired token"))
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.ID)
		ctx.Next()
	}
}
