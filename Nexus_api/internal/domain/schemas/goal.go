package schemas

import (
	"time"

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
	Status      GoalStatus      `gorm:"type:goal_status;not null;default:'Pendente'"`
	UserID      uint            `gorm:"not null"`
	User        User            `gorm:"foreignKey:UserID"`
	Tasks       []Task
}

func NewGoal(goal models.GoalName, description string, dates dates.GoalDates, userID uint) (*Goal, []*models.ErrorMessage) {
	if userID == 0 {
		return nil, []*models.ErrorMessage{
			models.NewErrorMessage("UserID", "is required"),
		}
	}

	return &Goal{
		GoalName:    goal,
		Description: description,
		Dates:       dates,
		Status:      GoalPending,
		UserID:      userID,
	}, nil
}

func (g *Goal) GetName() string {
	return g.GoalName.GetValue()
}

func (g *Goal) GetDescription() string {
	return g.Description
}

func (g *Goal) GetStatus() GoalStatus {
	return g.Status
}

func (g *Goal) GetStartDate() time.Time {
	return g.Dates.GetStartDateValue()
}

func (g *Goal) GetFinalDate() time.Time {
	return g.Dates.GetFinalDateValue()
}

func (g *Goal) GetTaskIDs() []uint {
	taskIDs := make([]uint, 0, len(g.Tasks))

	for _, task := range g.Tasks {
		taskIDs = append(taskIDs, task.ID)
	}

	return taskIDs
}

func (g *Goal) GetTasks() []Task {
	return g.Tasks
}

func (g *Goal) ChangeName(goalName models.GoalName) {
	g.GoalName = goalName
}

func (g *Goal) ChangeDescription(description string) {
	g.Description = description
}

func (g *Goal) ChangeStatus(status GoalStatus) {
	g.Status = status
}

func (g *Goal) ChangeDates(goalDates dates.GoalDates) {
	g.Dates = goalDates
}
