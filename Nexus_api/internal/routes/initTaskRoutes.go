package routes

import (
	th "github.com/MatheusMikio/Nexus/internal/handler"
	"github.com/MatheusMikio/Nexus/internal/service"
	"github.com/gin-gonic/gin"
)

func initDefaultTaskRoutes(rg *gin.RouterGroup, service service.ITaskService) {
	task := rg.Group("/goals/:goalId/tasks")
	{
		task.GET("", th.GetAllTasks(service))
		task.GET("/:taskId", th.GetTaskById(service))
		task.POST("", th.CreateTask(service))
		task.PUT("/:taskId", th.UpdateTask(service))
		task.DELETE("/:taskId", th.DeleteTask(service))
	}
}
