package routes

import (
	"github.com/MatheusMikio/Nexus/internal/container"
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"github.com/MatheusMikio/Nexus/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine, container *container.Container) {
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		publicRoutes := v1.Group("")
		{
			initDocsRoutes(publicRoutes)
			initPublicUserRoutes(publicRoutes, container.UserService)
		}

		v1.Use(middlewares.AuthMiddleware())

		defaultRoutes := v1.Group("")
		defaultRoutes.Use(middlewares.RoleMiddleware(schemas.Default, schemas.Admin))
		{
			initDefaultUserRoutes(defaultRoutes, container.UserService)
			initDefaultGoalRoutes(defaultRoutes, container.GoalService)
			initDefaultTaskRoutes(defaultRoutes, container.TaskService)
		}

		adminRoutes := v1.Group("")
		adminRoutes.Use(middlewares.RoleMiddleware(schemas.Admin))
		{
			initAdminUserRoutes(adminRoutes, container.UserService)
		}
	}
}
