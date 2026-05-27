package container

import (
	"gorm.io/gorm"
	"Nexus_api/internal/domain/repository"
	"Nexus_api/internal/domain/service"
)

type Container struct {
	GoalService IGoalService
	TaskService ITaskService
	UserService IUserService
}

func NewContainer(db *gorm.DB) *Container {
	goalRepo := repository.NewGoalRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)
	goalService := service.NewGoalService(goalRepo, userRepo)
	taskService := service.NewTaskService(taskRepo, userRepo)
	

	return &Container{
		GoalService: goalService,
		TaskService: taskService,
		UserService: userService,
	}
}
