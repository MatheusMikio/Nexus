package dates

import (
	"time"

	"github.com/MatheusMikio/Nexus/internal/domain/models"
)

type GoalDates struct {
	StartDate            time.Time `gorm:"not null"`
	FinalizationForecast time.Time
}

func NewGoalDates(startDate time.Time, finalizationForecast time.Time) (GoalDates, []*models.ErrorMessage) {
	return newGoalDates(startDate, finalizationForecast, true)
}

func NewGoalDatesFromExistingStart(startDate time.Time, finalizationForecast time.Time) (GoalDates, []*models.ErrorMessage) {
	return newGoalDates(startDate, finalizationForecast, false)
}

func newGoalDates(startDate time.Time, finalizationForecast time.Time, validateCurrentDate bool) (GoalDates, []*models.ErrorMessage) {
	errors := make([]*models.ErrorMessage, 0)

	validationLocation := startDate.Location()

	if startDate.IsZero() {
		errors = append(errors, models.NewErrorMessage("StartDate", "is required"))
	}

	if finalizationForecast.IsZero() {
		errors = append(errors, models.NewErrorMessage("FinalizationForecast", "is required"))
	}

	if validateCurrentDate && !startDate.IsZero() && CompareDate(startDate, time.Now(), validationLocation) < 0 {
		errors = append(errors, models.NewErrorMessage("StartDate", "must be greater than or equal to current date"))
	}

	if !startDate.IsZero() && !finalizationForecast.IsZero() && CompareDate(finalizationForecast, startDate, validationLocation) < 0 {
		errors = append(errors, models.NewErrorMessage("FinalizationForecast", "must be greater than or equal to start date"))
	}

	if len(errors) > 0 {
		return GoalDates{}, errors
	}

	return GoalDates{
		StartDate:            startDate,
		FinalizationForecast: finalizationForecast,
	}, nil
}

func (g *GoalDates) GetStartDateValue() time.Time {
	return g.StartDate
}

func (g *GoalDates) GetFinalDateValue() time.Time {
	return g.FinalizationForecast
}
