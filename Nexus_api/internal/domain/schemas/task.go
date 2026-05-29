package schemas

import (
	"time"

	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/dates"
	"gorm.io/gorm"
)

type TaskStatus string

const (
	TaskPending    TaskStatus = "Pendente"
	TaskInProgress TaskStatus = "Em progresso"
	TaskCompleted  TaskStatus = "Concluido"
)

type Task struct {
	gorm.Model
	Name        models.GoalName `gorm:"embedded"`
	Description string          `gorm:"default:'Descrição não informada'"`
	Status      TaskStatus      `gorm:"type:task_status;not null;default:'Pendente'"`
	Dates       dates.TaskDates `gorm:"embedded"`
	GoalID      uint            `gorm:"not null"`
	Goal        Goal            `gorm:"foreignKey:GoalID"`
}

func NewTask(name models.GoalName, description string, dates dates.TaskDates, goalID uint) (*Task, []*models.ErrorMessage) {
	if goalID == 0 {
		return nil, []*models.ErrorMessage{
			models.NewErrorMessage("GoalID", "is required"),
		}
	}

	if dates.GetFinalizationDate() != nil || dates.GetTimeSpent() != nil {
		return nil, []*models.ErrorMessage{
			models.NewErrorMessage("FinalizationDate", "must be empty when status is not completed"),
		}
	}

	return &Task{
		Name:        name,
		Description: description,
		Dates:       dates,
		GoalID:      goalID,
		Status:      TaskPending,
	}, nil
}

func (t *Task) GetName() string {
	return t.Name.GetValue()
}

func (t *Task) GetDescription() string {
	return t.Description
}

func (t *Task) GetStatus() TaskStatus {
	return t.Status
}

func (t *Task) GetStartDate() *time.Time {
	return t.Dates.GetStartDate()
}

func (t *Task) GetFinalizationDate() *time.Time {
	return t.Dates.GetFinalizationDate()
}

func (t *Task) GetTimeSpent() *int64 {
	return t.Dates.GetTimeSpent()
}

func (t *Task) GetGoalID() uint {
	return t.GoalID
}

func (t *Task) ChangeName(name models.GoalName) {
	t.Name = name
}

func (t *Task) ChangeDescription(description string) {
	t.Description = description
}

func (t *Task) ChangeStatus(status TaskStatus) {
	t.Status = status
}

func (t *Task) ChangeDates(taskDates dates.TaskDates) {
	t.Dates = taskDates
}
