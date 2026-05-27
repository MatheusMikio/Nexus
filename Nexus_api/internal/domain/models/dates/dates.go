package dates

import (
	"time"

	"github.com/MatheusMikio/Nexus/internal/domain/models"
)

type GoalDates struct {
	StartDate            time.Time `gorm:"not null"`
	FinalizationForecast *time.Time
}

type TaskDates struct {
	StartDate        *time.Time 
	FinalizationDate *time.Time
	TimeSpent        *int64
}

func NewGoalDates(startDate time.Time, finalizationForecast *time.Time) (GoalDates, []*models.ErrorMessage) {
	return GoalDates{
		StartDate:            startDate,
		FinalizationForecast: finalizationForecast,
	}, nil
}

func NewTaskDates(startDate time.Time, finalizationDate *time.Time) (TaskDates, []*models.ErrorMessage) {
	var timeSpent *int64

	if finalizationDate != nil {
		ts := int64(finalizationDate.Sub(startDate).Minutes())
		timeSpent = &ts
	}

	return TaskDates{
		StartDate:        startDate,
		FinalizationDate: finalizationDate,
		TimeSpent:        timeSpent,
	}, nil
}
