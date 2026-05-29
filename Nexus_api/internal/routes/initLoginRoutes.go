package routes

import (
	"github.com/MatheusMikio/Nexus/internal/handler"
	"github.com/MatheusMikio/Nexus/internal/service"
	"github.com/gin-gonic/gin"
)

func initLoginRoutes(rg *gin.RouterGroup, service service.ILoginService) {
	login := rg.Group("auth")
	{
		login.POST("", handler.LoginHandler(service))
	}
}
