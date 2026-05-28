package middlewares

import (
	"net/http"

	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"github.com/MatheusMikio/Nexus/internal/handler"
	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...schemas.Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, err := GetUserRole(ctx)
		if err != nil {
			handler.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", err.Error()))
			ctx.Abort()
			return
		}

		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				ctx.Next()
				return
			}
		}

		handler.SendError(ctx, http.StatusForbidden, models.NewErrorMessage("Authorization", "insufficient permissions"))
		ctx.Abort()
	}
}
