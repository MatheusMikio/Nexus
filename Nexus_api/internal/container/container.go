package container

import (
	"github.com/MatheusMikio/Nexus/internal/repository"
	"github.com/MatheusMikio/Nexus/internal/service"
	"gorm.io/gorm"
)

type Container struct {
	GoalService service.IGoalService
	TaskService service.ITaskService
	UserService service.IUserService
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
