package handler

import (
	"net/http"

	userdto "github.com/MatheusMikio/Nexus/internal/domain/dtos/user"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/helper"
	"github.com/MatheusMikio/Nexus/internal/response"
	"github.com/MatheusMikio/Nexus/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAllUsers godoc
// @Summary Listar usuarios
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Param page query int false "Pagina"
// @Param size query int false "Tamanho da pagina"
// @Success 200 {array} userdto.Response
// @Failure 400 {object} models.ErrorMessage
// @Failure 401 {object} models.ErrorMessage
// @Failure 403 {object} models.ErrorMessage
// @Failure 500 {object} models.ErrorMessage
// @Router /user [get]
func GetAllUsers(userService service.IUserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params, err := helper.BindPaginationQuery(ctx)
		if err != nil {
			response.SendError(ctx, http.StatusBadRequest, err)
			return
		}

		users, err := userService.GetAllUsers(params)
		if err != nil {
			response.SendError(ctx, http.StatusInternalServerError, err)
			return
		}

		response.SendSuccess(ctx, http.StatusOK, users)
	}
}

// GetUserById godoc
// @Summary Buscar usuario por ID
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID publico do usuario"
// @Success 200 {object} userdto.Response
// @Failure 400 {object} models.ErrorMessage
// @Failure 401 {object} models.ErrorMessage
// @Failure 403 {object} models.ErrorMessage
// @Failure 404 {object} models.ErrorMessage
// @Failure 500 {object} models.ErrorMessage
// @Router /user/{id} [get]
func GetUserById(userService service.IUserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := uuid.Parse(idStr)

		if err != nil {
			response.SendError(ctx, http.StatusBadRequest, models.NewErrorMessage("Validation", "Invalid user ID"))
			return
		}

		if !helper.AuthorizeSelfOrAdmin(ctx, id) {
			return
		}

		user, serviceErr := userService.GetUserById(id)

		if serviceErr != nil {
			response.SendError(ctx, helper.ErrorStatusCode(serviceErr), serviceErr)
			return
		}

		response.SendSuccess(ctx, http.StatusOK, user)
	}
}

// CreateUser godoc
// @Summary Criar usuario
// @Tags Users
// @Accept json
// @Produce json
// @Param request body userdto.Request true "Dados do usuario"
// @Success 201
// @Failure 400 {object} models.ErrorMessage
// @Failure 500 {object} models.ErrorMessage
// @Router /user [post]
func CreateUser(userService service.IUserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &userdto.Request{}

		if err := ctx.ShouldBindJSON(request); err != nil {
			response.SendError(ctx, http.StatusBadRequest, models.NewErrorMessage("Validation", err.Error()))
			return
		}

		if err := userService.CreateUser(request); err != nil {
			statusCode := http.StatusBadRequest
			if len(err) == 1 {
				statusCode = helper.ErrorStatusCode(err[0])
			}

			response.SendErrors(ctx, statusCode, err)
			return
		}

		response.SendSuccess(ctx, http.StatusCreated, nil)
	}
}

// UpdateUser godoc
// @Summary Atualizar usuario
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID publico do usuario"
// @Param request body userdto.Update true "Dados do usuario"
// @Success 200
// @Failure 400 {object} models.ErrorMessage
// @Failure 401 {object} models.ErrorMessage
// @Failure 403 {object} models.ErrorMessage
// @Failure 404 {object} models.ErrorMessage
// @Failure 500 {object} models.ErrorMessage
// @Router /user/{id} [put]
func UpdateUser(userService service.IUserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			response.SendError(ctx, http.StatusBadRequest, models.NewErrorMessage("Validation", "Invalid user ID"))
			return
		}

		if !helper.AuthorizeSelfOrAdmin(ctx, id) {
			return
		}

		request := &userdto.Update{}
		if err := ctx.ShouldBindJSON(request); err != nil {
			response.SendError(ctx, http.StatusBadRequest, models.NewErrorMessage("Validation", err.Error()))
			return
		}

		if err := userService.UpdateUser(id, request); err != nil {
			statusCode := http.StatusBadRequest
			if len(err) == 1 {
				statusCode = helper.ErrorStatusCode(err[0])
			}

			response.SendErrors(ctx, statusCode, err)
			return
		}

		response.SendSuccess(ctx, http.StatusOK, nil)
	}
}

// DeleteUser godoc
// @Summary Remover usuario
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID publico do usuario"
// @Success 200
// @Failure 400 {object} models.ErrorMessage
// @Failure 401 {object} models.ErrorMessage
// @Failure 403 {object} models.ErrorMessage
// @Failure 404 {object} models.ErrorMessage
// @Failure 500 {object} models.ErrorMessage
// @Router /user/{id} [delete]
func DeleteUser(userService service.IUserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		id, err := uuid.Parse(idStr)
		if err != nil {
			response.SendError(ctx, http.StatusBadRequest, models.NewErrorMessage("Validation", "Invalid user ID"))
			return
		}

		if !helper.AuthorizeSelfOrAdmin(ctx, id) {
			return
		}

		if err := userService.DeleteUser(id); err != nil {
			response.SendError(ctx, helper.ErrorStatusCode(err), err)
			return
		}

		response.SendSuccess(ctx, http.StatusOK, nil)
	}
}
