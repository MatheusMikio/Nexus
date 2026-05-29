package container

import (
	"github.com/MatheusMikio/Nexus/internal/repository"
	"github.com/MatheusMikio/Nexus/internal/service"
	"gorm.io/gorm"
)

type Container struct {
	GoalService  service.IGoalService
	TaskService  service.ITaskService
	UserService  service.IUserService
	LoginService service.ILoginService
}

func NewContainer(db *gorm.DB) *Container {
	goalRepo := repository.NewGoalRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	userRepo := repository.NewUserRepository(db)

	return &Container{
		GoalService:  service.NewGoalService(goalRepo, userRepo),
		TaskService:  service.NewTaskService(taskRepo, goalRepo, userRepo),
		UserService:  service.NewUserService(userRepo),
		LoginService: service.NewLoginService(userRepo),
	}
}
