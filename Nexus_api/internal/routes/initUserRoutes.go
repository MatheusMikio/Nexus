package routes

import (
	"github.com/MatheusMikio/Nexus/internal/handler"
	"github.com/MatheusMikio/Nexus/internal/service"
	"github.com/gin-gonic/gin"
)

func initPublicUserRoutes(router *gin.RouterGroup, userService service.IUserService) {
	user := router.Group("/user")
	{
		user.POST("", handler.Create(userService))
	}
}

func initDefaultUserRoutes(router *gin.RouterGroup, userService service.IUserService) {
	user := router.Group("/user")
	{
		user.GET("/:id", handler.GetUserById(userService))
		user.PUT("/:id", handler.Update(userService))
		user.DELETE("/:id", handler.Delete(userService))
	}
}

func initAdminUserRoutes(router *gin.RouterGroup, userService service.IUserService) {
	user := router.Group("/user")
	{
		user.GET("", handler.GetAllUsers(userService))
	}
}