package repository

import (
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"github.com/MatheusMikio/Nexus/internal/repository/base"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	base.ICrudRepository[schemas.Task]
	GetAllByGoalID(page, size int, goalID uint, userID uint) ([]*schemas.Task, error)
}

type TaskRepository struct {
	base.CrudRepository[schemas.Task]
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &TaskRepository{
		CrudRepository: base.CrudRepository[schemas.Task]{
			Db: db,
		},
	}
}

func (tr *TaskRepository) GetAllByGoalID(page, size int, goalID uint, userID uint) ([]*schemas.Task, error) {
	var tasks []*schemas.Task
	offset := (page - 1) * size

	err := tr.Db.
		Joins("JOIN goals ON goals.id = tasks.goal_id").
		Where("tasks.goal_id = ? AND goals.user_id = ?", goalID, userID).
		Offset(offset).
		Limit(size).
		Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
