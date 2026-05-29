package dates

import (
	"time"

	"github.com/MatheusMikio/Nexus/internal/domain/models"
)

type TaskDates struct {
	StartDate        *time.Time
	FinalizationDate *time.Time
	TimeSpent        *int64
}

func NewTaskDates(startDate *time.Time, finalizationDate *time.Time) (TaskDates, []*models.ErrorMessage) {
	errors := make([]*models.ErrorMessage, 0)
	var timeSpent *int64

	if startDate == nil && finalizationDate != nil {
		errors = append(errors, models.NewErrorMessage("StartDate", "is required when finalization date is informed"))
	}

	if startDate != nil && finalizationDate != nil && !finalizationDate.After(*startDate) {
		errors = append(errors, models.NewErrorMessage("FinalizationDate", "must be greater than start date"))
	}

	if len(errors) > 0 {
		return TaskDates{}, errors
	}

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
