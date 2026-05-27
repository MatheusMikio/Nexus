package container

import (
	"gorm.io/gorm"
)

type Container struct {
}

func NewContainer(db *gorm.DB) *Container {
	// goalRepo := repository.NewGoalRepository(db)
	// taskRepo := repository.NewTaskRepository(db)
	// userRepo := repository.NewUserRepository(db)

	return &Container{}
}
