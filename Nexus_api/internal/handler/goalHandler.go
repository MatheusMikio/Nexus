package handler

import (
	"net/http"

	"github.com/MatheusMikio/Nexus/internal/domain/dtos/goal"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/helper"
	"github.com/MatheusMikio/Nexus/internal/middlewares"
	"github.com/MatheusMikio/Nexus/internal/response"
	"github.com/MatheusMikio/Nexus/internal/service"
	"github.com/gin-gonic/gin"
)

// GetAllGoals godoc
// @Summary Listar metas do usuario autenticado
// @Tags Goals
// @Produce json
// @Security BearerAuth
// @Param page query int false "Pagina"
// @Param size query int false "Tamanho da pagina"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /goals [get]
func GetAllGoals(service service.IGoalService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := middlewares.GetUserID(ctx)
		if err != nil {
			response.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "unauthorized"))
			return
		}

		params, errBinding := helper.BindPaginationQuery(ctx)
		if errBinding != nil {
			response.SendError(ctx, http.StatusBadRequest, errBinding)
			return
		}

		goals, errService := service.GetAllGoals(params, userId)
		if errService != nil {
			response.SendError(ctx, helper.ErrorStatusCode(errService), errService)
			return
		}

		response.SendSuccess(ctx, http.StatusOK, goals)
	}
}

// GetGoalById godoc
// @Summary Buscar meta por ID
// @Tags Goals
// @Produce json
// @Security BearerAuth
// @Param goalId path int true "ID da meta"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /goals/{goalId} [get]
func GetGoalById(service service.IGoalService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := middlewares.GetUserID(ctx)
		if err != nil {
			response.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "unauthorized"))
			return
		}

		id, ok := helper.ParseUintParam(ctx, "goalId")
		if !ok {
			return
		}

		result, errService := service.GetById(id, userId)
		if errService != nil {
			response.SendError(ctx, helper.ErrorStatusCode(errService), errService)
			return
		}

		response.SendSuccess(ctx, http.StatusOK, result)

	}
}

// CreateGoal godoc
// @Summary Criar meta
// @Tags Goals
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body goal.Request true "Dados da meta"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /goals [post]
func CreateGoal(service service.IGoalService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := middlewares.GetUserID(ctx)
		if err != nil {
			response.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "unauthorized"))
			return
		}

		request := &goal.Request{}
		if err := ctx.ShouldBindJSON(request); err != nil {
			response.SendError(ctx, http.StatusBadRequest, models.NewErrorMessage("Validation", err.Error()))
			return
		}

		if err := service.Create(request, userId); err != nil {
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

// UpdateGoal godoc
// @Summary Atualizar meta
// @Tags Goals
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param goalId path int true "ID da meta"
// @Param request body goal.Update true "Dados da meta"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /goals/{goalId} [put]
func UpdateGoal(service service.IGoalService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := middlewares.GetUserID(ctx)
		if err != nil {
			response.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "unauthorized"))
			return
		}

		id, ok := helper.ParseUintParam(ctx, "goalId")
		if !ok {
			return
		}

		request := &goal.Update{}
		if err := ctx.ShouldBindJSON(request); err != nil {
			response.SendError(ctx, http.StatusBadRequest, models.NewErrorMessage("Validation", err.Error()))
			return
		}

		if err := service.Update(id, request, userId); err != nil {
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

// DeleteGoal godoc
// @Summary Remover meta
// @Tags Goals
// @Produce json
// @Security BearerAuth
// @Param goalId path int true "ID da meta"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /goals/{goalId} [delete]
func DeleteGoal(service service.IGoalService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := middlewares.GetUserID(ctx)
		if err != nil {
			response.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "unauthorized"))
			return
		}

		id, ok := helper.ParseUintParam(ctx, "goalId")
		if !ok {
			return
		}

		if err := service.Delete(id, userId); err != nil {
			response.SendError(ctx, helper.ErrorStatusCode(err), err)
			return
		}

		response.SendSuccess(ctx, http.StatusOK, nil)

	}
}
