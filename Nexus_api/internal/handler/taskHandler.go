package handler

import (
	"net/http"

	"github.com/MatheusMikio/Nexus/internal/domain/dtos/task"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/helper"
	"github.com/MatheusMikio/Nexus/internal/middlewares"
	"github.com/MatheusMikio/Nexus/internal/response"
	"github.com/MatheusMikio/Nexus/internal/service"
	"github.com/gin-gonic/gin"
)

// GetAllTasks godoc
// @Summary Listar tarefas de uma meta
// @Tags Tasks
// @Produce json
// @Security BearerAuth
// @Param goalId path int true "ID da meta"
// @Param page query int false "Pagina"
// @Param size query int false "Tamanho da pagina"
// @Success 200 {array} task.Response
// @Failure 400 {object} models.ErrorMessage
// @Failure 401 {object} models.ErrorMessage
// @Failure 404 {object} models.ErrorMessage
// @Failure 500 {object} models.ErrorMessage
// @Router /goals/{goalId}/tasks [get]
func GetAllTasks(service service.ITaskService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := middlewares.GetUserID(ctx)
		if err != nil {
			response.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "unauthorized"))
			return
		}

		goalId, ok := helper.ParseUintParam(ctx, "goalId")
		if !ok {
			return
		}

		params, errBinding := helper.BindPaginationQuery(ctx)
		if errBinding != nil {
			response.SendError(ctx, http.StatusBadRequest, errBinding)
			return
		}

		tasks, errService := service.GetAllTasks(params, goalId, userId)
		if errService != nil {
			response.SendError(ctx, helper.ErrorStatusCode(errService), errService)
			return
		}

		response.SendSuccess(ctx, http.StatusOK, tasks)
	}
}

// GetTaskById godoc
// @Summary Buscar tarefa por ID
// @Tags Tasks
// @Produce json
// @Security BearerAuth
// @Param goalId path int true "ID da meta"
// @Param taskId path int true "ID da tarefa"
// @Success 200 {object} task.Response
// @Failure 400 {object} models.ErrorMessage
// @Failure 401 {object} models.ErrorMessage
// @Failure 404 {object} models.ErrorMessage
// @Failure 500 {object} models.ErrorMessage
// @Router /goals/{goalId}/tasks/{taskId} [get]
func GetTaskById(service service.ITaskService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := middlewares.GetUserID(ctx)
		if err != nil {
			response.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "unauthorized"))
			return
		}

		goalId, ok := helper.ParseUintParam(ctx, "goalId")
		if !ok {
			return
		}

		taskId, ok := helper.ParseUintParam(ctx, "taskId")
		if !ok {
			return
		}

		result, errService := service.GetById(goalId, taskId, userId)
		if errService != nil {
			response.SendError(ctx, helper.ErrorStatusCode(errService), errService)
			return
		}

		response.SendSuccess(ctx, http.StatusOK, result)
	}
}

// CreateTask godoc
// @Summary Criar tarefa
// @Tags Tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param goalId path int true "ID da meta"
// @Param request body task.Request true "Dados da tarefa"
// @Success 201
// @Failure 400 {object} models.ErrorMessage
// @Failure 401 {object} models.ErrorMessage
// @Failure 404 {object} models.ErrorMessage
// @Failure 500 {object} models.ErrorMessage
// @Router /goals/{goalId}/tasks [post]
func CreateTask(service service.ITaskService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := middlewares.GetUserID(ctx)
		if err != nil {
			response.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "unauthorized"))
			return
		}

		goalId, ok := helper.ParseUintParam(ctx, "goalId")
		if !ok {
			return
		}

		request := &task.Request{}
		if err := ctx.ShouldBindJSON(request); err != nil {
			response.SendError(ctx, http.StatusBadRequest, models.NewErrorMessage("Validation", err.Error()))
			return
		}

		if err := service.Create(goalId, request, userId); err != nil {
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

// UpdateTask godoc
// @Summary Atualizar tarefa
// @Tags Tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param goalId path int true "ID da meta"
// @Param taskId path int true "ID da tarefa"
// @Param request body task.Update true "Dados da tarefa"
// @Success 200
// @Failure 400 {object} models.ErrorMessage
// @Failure 401 {object} models.ErrorMessage
// @Failure 404 {object} models.ErrorMessage
// @Failure 500 {object} models.ErrorMessage
// @Router /goals/{goalId}/tasks/{taskId} [put]
func UpdateTask(service service.ITaskService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := middlewares.GetUserID(ctx)
		if err != nil {
			response.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "unauthorized"))
			return
		}

		goalId, ok := helper.ParseUintParam(ctx, "goalId")
		if !ok {
			return
		}

		taskId, ok := helper.ParseUintParam(ctx, "taskId")
		if !ok {
			return
		}

		request := &task.Update{}
		if err := ctx.ShouldBindJSON(request); err != nil {
			response.SendError(ctx, http.StatusBadRequest, models.NewErrorMessage("Validation", err.Error()))
			return
		}

		if err := service.Update(goalId, taskId, request, userId); err != nil {
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

// DeleteTask godoc
// @Summary Remover tarefa
// @Tags Tasks
// @Produce json
// @Security BearerAuth
// @Param goalId path int true "ID da meta"
// @Param taskId path int true "ID da tarefa"
// @Success 200
// @Failure 400 {object} models.ErrorMessage
// @Failure 401 {object} models.ErrorMessage
// @Failure 404 {object} models.ErrorMessage
// @Failure 500 {object} models.ErrorMessage
// @Router /goals/{goalId}/tasks/{taskId} [delete]
func DeleteTask(service service.ITaskService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := middlewares.GetUserID(ctx)
		if err != nil {
			response.SendError(ctx, http.StatusUnauthorized, models.NewErrorMessage("Authorization", "unauthorized"))
			return
		}

		goalId, ok := helper.ParseUintParam(ctx, "goalId")
		if !ok {
			return
		}

		taskId, ok := helper.ParseUintParam(ctx, "taskId")
		if !ok {
			return
		}

		if err := service.Delete(goalId, taskId, userId); err != nil {
			response.SendError(ctx, helper.ErrorStatusCode(err), err)
			return
		}

		response.SendSuccess(ctx, http.StatusOK, nil)
	}
}
