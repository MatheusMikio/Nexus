package routes

func initUserRoutes(router *gin.RouterGroup, userService service.IUserService){
	user := router.Group("/user")
	{
		user.GET("", handler.GetAll(userService))
		user.GET("/:id", handler.GetById(userService))
		user.POST("", handler.Create(userService))
		user.PUT("/:id", handler.Update(userService))
		user.DELETE("/:id", handler.Delete(userService))
	}
}