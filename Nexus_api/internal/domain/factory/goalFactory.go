package factory

import (
	"time"

	dto "github.com/MatheusMikio/Nexus/internal/domain/dtos/goal"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/dates"
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"github.com/MatheusMikio/Nexus/internal/helper"
)

func NewGoalFromRequest(gr *dto.Request, userID uint) (*schemas.Goal, []*models.ErrorMessage) {
	errors := make([]*models.ErrorMessage, 0)

	goalName, err := models.NewGoalName(gr.Name)
	errors = helper.AppendErrors(errors, err)

	goalDates, err := dates.NewGoalDates(gr.StartDate, gr.EndDate)
	errors = helper.AppendErrors(errors, err)

	if len(errors) > 0 {
		return nil, errors
	}

	return schemas.NewGoal(goalName, gr.Description, goalDates, userID)
}

func BuildGoalUpdate(gr *dto.Update, goalDb *schemas.Goal) []*models.ErrorMessage {
	errors := make([]*models.ErrorMessage, 0)

	goalName, err := buildGoalName(gr.Name)
	errors = helper.AppendErrors(errors, err)

	goalStatus, errMessage := buildGoalStatus(gr.Status)
	if errMessage != nil {
		errors = append(errors, errMessage)
	}

	startDateChanged := hasGoalStartDateChanged(goalDb.GetStartDate(), gr.StartDate)

	goalDates, err := buildGoalDates(gr, goalDb, startDateChanged)
	errors = helper.AppendErrors(errors, err)

	if startDateChanged {
		errors = appendGoalTaskStartDateErrors(errors, *gr.StartDate, goalDb)
	}

	if len(errors) > 0 {
		return errors
	}

	if goalName != nil {
		goalDb.ChangeName(*goalName)
	}

	if gr.Description != nil {
		goalDb.ChangeDescription(*gr.Description)
	}

	if goalStatus != nil {
		goalDb.ChangeStatus(*goalStatus)
	}

	if goalDates != nil {
		goalDb.ChangeDates(*goalDates)
	}

	return nil
}

func buildGoalName(value *string) (*models.GoalName, []*models.ErrorMessage) {
	if value == nil {
		return nil, nil
	}

	goalName, errors := models.NewGoalName(*value)
	if len(errors) > 0 {
		return nil, errors
	}

	return &goalName, nil
}

func buildGoalStatus(value *string) (*schemas.GoalStatus, *models.ErrorMessage) {
	if value == nil {
		return nil, nil
	}

	switch schemas.GoalStatus(*value) {
	case schemas.GoalPending, schemas.GoalCompleted, schemas.GoalCanceled, schemas.GoalLate:
		status := schemas.GoalStatus(*value)
		return &status, nil
	default:
		return nil, models.NewErrorMessage("Status", "invalid")
	}
}

func buildGoalDates(gr *dto.Update, goalDb *schemas.Goal, startDateChanged bool) (*dates.GoalDates, []*models.ErrorMessage) {
	if gr.StartDate == nil && gr.EndDate == nil {
		return nil, nil
	}

	startDate := goalDb.GetStartDate()
	if gr.StartDate != nil {
		startDate = *gr.StartDate
	}

	finalizationForecast := goalDb.GetFinalDate()
	if gr.EndDate != nil {
		finalizationForecast = *gr.EndDate
	}

	goalDates, errors := buildGoalDatesValue(startDate, finalizationForecast, startDateChanged)
	if len(errors) > 0 {
		return nil, errors
	}

	return &goalDates, nil
}

func buildGoalDatesValue(startDate, finalizationForecast time.Time, startDateChanged bool) (dates.GoalDates, []*models.ErrorMessage) {
	if startDateChanged {
		return dates.NewGoalDates(startDate, finalizationForecast)
	}

	return dates.NewGoalDatesFromExistingStart(startDate, finalizationForecast)
}

func hasGoalStartDateChanged(current time.Time, next *time.Time) bool {
	if next == nil {
		return false
	}

	return !current.Equal(*next)
}

func appendGoalTaskStartDateErrors(errors []*models.ErrorMessage, startDate time.Time, goalDb *schemas.Goal) []*models.ErrorMessage {
	startDateOnly := dateOnly(startDate)

	for _, task := range goalDb.GetTasks() {
		taskStartDate := task.GetStartDate()
		if taskStartDate == nil {
			continue
		}

		if dateOnly(*taskStartDate).Before(startDateOnly) {
			return append(errors, models.NewErrorMessage("StartDate", "must be less than or equal to existing task start dates"))
		}
	}

	return errors
}
