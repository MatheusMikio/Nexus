package routes

import (
	"github.com/MatheusMikio/Nexus/internal/handler"
	"github.com/MatheusMikio/Nexus/internal/service"
	"github.com/gin-gonic/gin"
)

func initPublicUserRoutes(router *gin.RouterGroup, service service.IUserService) {
	user := router.Group("/user")
	{
		user.POST("", handler.CreateUser(service))
	}
}

func initDefaultUserRoutes(router *gin.RouterGroup, service service.IUserService) {
	user := router.Group("/user")
	{
		user.GET("/:id", handler.GetUserById(service))
		user.PUT("/:id", handler.UpdateUser(service))
		user.DELETE("/:id", handler.DeleteUser(service))
	}
}

func initAdminUserRoutes(router *gin.RouterGroup, service service.IUserService) {
	user := router.Group("/user")
	{
		user.GET("", handler.GetAllUsers(service))
	}
}
