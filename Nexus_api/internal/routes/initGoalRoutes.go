package routes

import (
	gh "github.com/MatheusMikio/Nexus/internal/handler"
	"github.com/MatheusMikio/Nexus/internal/service"
	"github.com/gin-gonic/gin"
)

func initDefaultGoalRoutes(rg *gin.RouterGroup, service service.IGoalService) {
	goal := rg.Group("/goals")
	{
		goal.GET("", gh.GetAllGoals(service))
		goal.GET("/:goalId", gh.GetGoalById(service))
		goal.POST("", gh.CreateGoal(service))
		goal.PUT("/:goalId", gh.UpdateGoal(service))
		goal.DELETE("/:goalId", gh.DeleteGoal(service))
	}
}
