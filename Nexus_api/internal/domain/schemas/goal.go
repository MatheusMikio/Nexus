package schemas

import (
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/dates"
	"gorm.io/gorm"
)

type GoalStatus string

const (
	GoalPending   GoalStatus = "Pendente"
	GoalCompleted GoalStatus = "Concluido"
	GoalCanceled  GoalStatus = "Cancelado"
	GoalLate      GoalStatus = "Atrasada"
)

type Goal struct {
	gorm.Model
	GoalName    models.GoalName `gorm:"embedded"`
	Description string          `gorm:"default:'Descrição não informada'"`
	Dates       dates.GoalDates `gorm:"embedded"`
	Status      GoalStatus      `gorm:"type:goal_status;not null;default:'Pending'"`
	UserID      uint            `gorm:"not null"`
	User        User            `gorm:"foreignKey:UserID"`
	Tasks       []Task
}

func NewGoal(goal models.GoalName, description string, dates dates.GoalDates, userID uint) (*Goal, []*models.ErrorMessage) {
	return &Goal{
		GoalName:    goal,
		Description: description,
		Dates: dates,
		Status: GoalPending,
		UserID: userID,
	}, nil
}