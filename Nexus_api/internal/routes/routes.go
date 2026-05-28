package routes

func initRoutes(router *gin.Engine, container *container.Container){
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		v1.Use(middleware.AuthMiddleware())
		initUserRoutes(v1, container.UserService)
		initGoalRoutes(v1, container.GoalService)
		initTaskRoutes(v1, container.TaskService)
	}
}