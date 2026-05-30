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
		user.GET("/me", handler.GetMe(service))
		user.PUT("/me", handler.UpdateMe(service))
		user.DELETE("/me", handler.DeleteMe(service))
	}
}

func initAdminUserRoutes(router *gin.RouterGroup, service service.IUserService) {
	user := router.Group("/user")
	{
		user.GET("", handler.GetAllUsers(service))
	}

	adminUser := router.Group("/admin/user")
	{
		adminUser.GET("/:id", handler.GetUserById(service))
		adminUser.PUT("/:id", handler.UpdateUser(service))
		adminUser.DELETE("/:id", handler.DeleteUser(service))
	}
}
