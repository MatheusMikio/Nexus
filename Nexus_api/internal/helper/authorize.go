package helper

import (
	"net/http"

	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"github.com/MatheusMikio/Nexus/internal/middlewares"
	"github.com/MatheusMikio/Nexus/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthorizeSelfOrAdmin(ctx *gin.Context, targetUserID uuid.UUID) bool {
	authenticatedUserID, err := middlewares.GetUserID(ctx)
	if err != nil {
		response.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "unauthorized"))
		return false
	}

	authenticatedUserRole, err := middlewares.GetUserRole(ctx)
	if err != nil {
		response.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "unauthorized"))
		return false
	}

	if authenticatedUserRole != schemas.Admin && authenticatedUserID != targetUserID {
		response.SendError(ctx, http.StatusForbidden, models.NewErrorMessage("Authorization", "insufficient permissions"))
		return false
	}

	return true
}
