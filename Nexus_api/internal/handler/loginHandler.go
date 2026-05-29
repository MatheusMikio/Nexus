package handler

import (
	"net/http"

	"github.com/MatheusMikio/Nexus/internal/domain/dtos/auth"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/response"
	"github.com/MatheusMikio/Nexus/internal/service"
	"github.com/gin-gonic/gin"
)

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

		response.SendAuthSuccess(ctx, http.StatusOK, user.AccessToken, user.ExpiresIn, user.User)
	}
}
