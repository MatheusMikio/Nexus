package handler

import (
	"net/http"

	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"github.com/MatheusMikio/Nexus/internal/helper"
	"github.com/MatheusMikio/Nexus/internal/middlewares"
	"github.com/MatheusMikio/Nexus/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllUsers(userService service.IUserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params, err := helper.BindPaginationQuery(ctx)
		if err != nil {
			SendError(ctx, http.StatusBadRequest, err)
			return
		}

		users, err := userService.GetAll(params)
		if err != nil {
			SendError(ctx, http.StatusInternalServerError, err)
			return
		}

		SendSuccess(ctx, http.StatusOK, users)
	}
}

func GetUserById(userService service.IUserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := uuid.Parse(idStr)

		if err != nil {
			SendError(ctx, http.StatusBadRequest, models.NewErrorMessage("Validation", "Invalid user ID"))
			return
		}

		if !authorizeSelfOrAdmin(ctx, id) {
			return
		}

		user, serviceErr := userService.GetById(id)
		
		if serviceErr != nil {
			SendError(ctx, http.StatusNotFound, serviceErr)
			return
		}

		SendSuccess(ctx, http.StatusOK, user)
	}
}

func authorizeSelfOrAdmin(ctx *gin.Context, targetUserID uuid.UUID) bool {
	authenticatedUserID, err := middlewares.GetUserID(ctx)
	if err != nil {
		SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "unauthorized"))
		return false
	}

	authenticatedUserRole, err := middlewares.GetUserRole(ctx)
	if err != nil {
		SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "unauthorized"))
		return false
	}

	if authenticatedUserRole != schemas.Admin && authenticatedUserID != targetUserID {
		SendError(ctx, http.StatusForbidden, models.NewErrorMessage("Authorization", "insufficient permissions"))
		return false
	}

	return true
}
