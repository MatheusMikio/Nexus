package handler

import (
	"net/http"

	userdto "github.com/MatheusMikio/Nexus/internal/domain/dtos/user"
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
			SendError(ctx, userErrorStatusCode(serviceErr), serviceErr)
			return
		}

		SendSuccess(ctx, http.StatusOK, user)
	}
}

func Create(userService service.IUserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &userdto.Request{}

		if err := ctx.ShouldBindJSON(request); err != nil {
			SendError(ctx, http.StatusBadRequest, models.NewErrorMessage("Validation", err.Error()))
			return
		}

		if err := userService.Create(request); err != nil {
			SendErrors(ctx, http.StatusBadRequest, err)
			return
		}

		SendSuccess(ctx, http.StatusCreated, nil)
	}
}

func Update(userService service.IUserService) gin.HandlerFunc {
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

		request := &userdto.Update{}
		if err := ctx.ShouldBindJSON(request); err != nil {
			SendError(ctx, http.StatusBadRequest, models.NewErrorMessage("Validation", err.Error()))
			return
		}

		if err := userService.Update(id, request); err != nil {
			statusCode := http.StatusBadRequest
			if len(err) == 1 {
				statusCode = userErrorStatusCode(err[0])
			}

			SendErrors(ctx, statusCode, err)
			return
		}

		SendSuccess(ctx, http.StatusOK, nil)
	}
}

func Delete(userService service.IUserService) gin.HandlerFunc {
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

		if err := userService.Delete(id); err != nil {
			SendError(ctx, userErrorStatusCode(err), err)
			return
		}

		SendSuccess(ctx, http.StatusOK, nil)
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

func userErrorStatusCode(err *models.ErrorMessage) int {
	if err.Property == "User" && err.Message == "Not found" {
		return http.StatusNotFound
	}

	if err.Property == "Database" {
		return http.StatusInternalServerError
	}

	return http.StatusBadRequest
}
