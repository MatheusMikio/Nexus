package handler

import (
	"net/http"

	"github.com/MatheusMikio/Nexus/internal/domain/dtos/auth"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/response"
	"github.com/MatheusMikio/Nexus/internal/service"
	"github.com/gin-gonic/gin"
)

// LoginHandler godoc
// @Summary Autenticar usuario
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body auth.Login true "Credenciais do usuario"
// @Success 200 {object} auth.LoginResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 422 {object} response.ErrorResponse
// @Router /auth [post]
func LoginHandler(service service.ILoginService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &auth.Login{}

		if err := ctx.ShouldBindJSON(request); err != nil {
			response.SendError(ctx, http.StatusBadRequest, models.NewErrorMessage("Validation", err.Error()))
			return
		}

		user, errService := service.Login(request.Email, request.Password)
		if errService != nil {
			response.SendError(ctx, http.StatusUnprocessableEntity, errService)
			return
		}

		response.SendAuthSuccess(ctx, http.StatusOK, user)
	}
}
