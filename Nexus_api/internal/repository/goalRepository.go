package repository

import (
	"errors"

	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"github.com/MatheusMikio/Nexus/internal/repository/base"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IGoalRepository interface {
	base.ICrudRepository[schemas.Goal]
	GetAllByUserId(page, size int, userId uuid.UUID) ([]*schemas.Goal, error)
	GetByIDAndUserID(goalID, userID uint) (*schemas.Goal, error)
}

type GoalRepository struct {
	base.CrudRepository[schemas.Goal]
}

func NewGoalRepository(db *gorm.DB) IGoalRepository {
	return &GoalRepository{
		CrudRepository: base.CrudRepository[schemas.Goal]{
			Db: db,
		},
	}
}

func (gr *GoalRepository) GetAllByUserId(page, size int, userId uuid.UUID) ([]*schemas.Goal, error) {
	var goals []*schemas.Goal
	offset := (page - 1) * size

	if err := gr.Db.
		Preload("Tasks").
		Joins("JOIN users ON users.id = goals.user_id").
		Where("users.public_id = ?", userId).
		Offset(offset).
		Limit(size).
		Find(&goals).Error; err != nil {
		return nil, err
	}
	return goals, nil
}

func (r *GoalRepository) GetByIDAndUserID(goalID, userID uint) (*schemas.Goal, error) {
	var goal schemas.Goal

	err := r.Db.
		Preload("Tasks").
		Where("id = ? AND user_id = ?", goalID, userID).
		First(&goal).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, base.ErrNotFound
		}

		return nil, err
	}

	return &goal, nil
}
