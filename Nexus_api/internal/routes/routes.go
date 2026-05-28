package routes

import (
	"github.com/MatheusMikio/Nexus/internal/container"
	"github.com/MatheusMikio/Nexus/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine, container *container.Container) {
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		initPublicUserRoutes(v1, container.UserService)

		v1.Use(middlewares.AuthMiddleware())

		initUserRoutes(v1, container.UserService)
		initGoalRoutes(v1, container.GoalService)
		initTaskRoutes(v1, container.TaskService)
	}
}
