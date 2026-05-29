package dates

import (
	"time"

	"github.com/MatheusMikio/Nexus/internal/domain/models"
)

type GoalDates struct {
	StartDate            time.Time `gorm:"not null"`
	FinalizationForecast *time.Time
}

func NewGoalDates(startDate time.Time, finalizationForecast *time.Time) (GoalDates, []*models.ErrorMessage) {
	return GoalDates{
		StartDate:            startDate,
		FinalizationForecast: finalizationForecast,
	}, nil
}

func (g *GoalDates) GetStartDateValue() time.Time {
	return g.StartDate
}

func (g *GoalDates) GetFinalDateValue() *time.Time {
	return g.FinalizationForecast
}

type TaskDates struct {
	StartDate        *time.Time
	FinalizationDate *time.Time
	TimeSpent        *int64
}

func NewTaskDates(startDate *time.Time, finalizationDate *time.Time) (TaskDates, []*models.ErrorMessage) {
	var timeSpent *int64

	if startDate != nil && finalizationDate != nil {
		ts := int64(finalizationDate.Sub(*startDate).Minutes())
		timeSpent = &ts
	}

	return TaskDates{
		StartDate:        startDate,
		FinalizationDate: finalizationDate,
		TimeSpent:        timeSpent,
	}, nil
}

func (t *TaskDates) GetStartDate() *time.Time {
	return t.StartDate
}

func (t *TaskDates) GetFinalizationDate() *time.Time {
	return t.FinalizationDate
}

func (t *TaskDates) GetTimeSpent() *int64 {
	return t.TimeSpent
}
