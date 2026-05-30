package factory

import (
	"time"

	dto "github.com/MatheusMikio/Nexus/internal/domain/dtos/task"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/dates"
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"github.com/MatheusMikio/Nexus/internal/helper"
)

func NewTaskFromRequest(tr *dto.Request, goal *schemas.Goal) (*schemas.Task, []*models.ErrorMessage) {
	errors := make([]*models.ErrorMessage, 0)

	taskName, err := models.NewGoalName(tr.Name)
	errors = helper.AppendErrors(errors, err)

	errors = appendTaskStartDateError(errors, tr.StartDate)
	errors = appendTaskGoalStartDateError(errors, tr.StartDate, goal)

	taskDates, err := dates.NewTaskDates(tr.StartDate, nil)
	errors = helper.AppendErrors(errors, err)

	if len(errors) > 0 {
		return nil, errors
	}

	return schemas.NewTask(taskName, tr.Description, taskDates, goal.ID)
}

func BuildTaskUpdate(tr *dto.Update, taskDb *schemas.Task, goal *schemas.Goal) []*models.ErrorMessage {
	errors := make([]*models.ErrorMessage, 0)

	taskName, err := buildTaskName(tr.Name)
	errors = helper.AppendErrors(errors, err)

	taskStatus, errMessage := buildTaskStatus(tr.Status)
	if errMessage != nil {
		errors = append(errors, errMessage)
	}

	effectiveStatus := taskDb.GetStatus()
	if taskStatus != nil {
		effectiveStatus = *taskStatus
	}

	taskDates, err := buildTaskDates(tr, taskDb, goal, effectiveStatus)
	errors = helper.AppendErrors(errors, err)

	if len(errors) > 0 {
		return errors
	}

	if taskName != nil {
		taskDb.ChangeName(*taskName)
	}

	if tr.Description != nil {
		taskDb.ChangeDescription(*tr.Description)
	}

	if taskStatus != nil {
		taskDb.ChangeStatus(*taskStatus)
	}

	if taskDates != nil {
		taskDb.ChangeDates(*taskDates)
	}

	return nil
}

func buildTaskName(value *string) (*models.GoalName, []*models.ErrorMessage) {
	if value == nil {
		return nil, nil
	}

	taskName, errors := models.NewGoalName(*value)
	if len(errors) > 0 {
		return nil, errors
	}

	return &taskName, nil
}

func buildTaskStatus(value *string) (*schemas.TaskStatus, *models.ErrorMessage) {
	if value == nil {
		return nil, nil
	}

	switch schemas.TaskStatus(*value) {
	case schemas.TaskPending, schemas.TaskInProgress, schemas.TaskCompleted:
		status := schemas.TaskStatus(*value)
		return &status, nil
	default:
		return nil, models.NewErrorMessage("Status", "invalid")
	}
}

func buildTaskDates(tr *dto.Update, taskDb *schemas.Task, goal *schemas.Goal, status schemas.TaskStatus) (*dates.TaskDates, []*models.ErrorMessage) {
	shouldRebuildDates := tr.StartDate != nil || tr.FinalizationDate != nil || status != taskDb.GetStatus()
	if !shouldRebuildDates {
		return nil, nil
	}

	startDate := taskDb.GetStartDate()
	startDateChanged := false
	if tr.StartDate != nil {
		startDate = tr.StartDate
		startDateChanged = hasTaskStartDateChanged(taskDb.GetStartDate(), tr.StartDate)
	}

	errors := make([]*models.ErrorMessage, 0)
	if startDateChanged {
		errors = appendTaskStartDateError(errors, tr.StartDate)
		errors = appendTaskGoalStartDateError(errors, tr.StartDate, goal)
	}

	if status != schemas.TaskCompleted {
		if tr.FinalizationDate != nil {
			errors = append(errors, models.NewErrorMessage("FinalizationDate", "must be empty when status is not completed"))
		}

		taskDates, err := dates.NewTaskDates(startDate, nil)
		errors = helper.AppendErrors(errors, err)
		if len(errors) > 0 {
			return nil, errors
		}

		return &taskDates, nil
	}

	finalizationDate := taskDb.GetFinalizationDate()
	if tr.FinalizationDate != nil {
		finalizationDate = tr.FinalizationDate
	}

	if finalizationDate == nil {
		errors = append(errors, models.NewErrorMessage("FinalizationDate", "is required when status is completed"))
	}

	if startDate == nil {
		errors = append(errors, models.NewErrorMessage("StartDate", "is required when status is completed"))
	}

	taskDates, err := dates.NewTaskDates(startDate, finalizationDate)
	errors = helper.AppendErrors(errors, err)
	if len(errors) > 0 {
		return nil, errors
	}

	return &taskDates, nil
}

func hasTaskStartDateChanged(current *time.Time, next *time.Time) bool {
	if current == nil || next == nil {
		return current != next
	}

	return dates.CompareDate(*next, *current, current.Location()) != 0
}

func appendTaskStartDateError(errors []*models.ErrorMessage, startDate *time.Time) []*models.ErrorMessage {
	if startDate == nil {
		return errors
	}

	if dates.CompareDate(*startDate, time.Now(), startDate.Location()) < 0 {
		return append(errors, models.NewErrorMessage("StartDate", "must be greater than or equal to current date"))
	}

	return errors
}

func appendTaskGoalStartDateError(errors []*models.ErrorMessage, startDate *time.Time, goal *schemas.Goal) []*models.ErrorMessage {
	if startDate == nil {
		return errors
	}

	goalStartDate := goal.GetStartDate()
	if dates.CompareDate(*startDate, goalStartDate, goalStartDate.Location()) < 0 {
		return append(errors, models.NewErrorMessage("StartDate", "must be greater than or equal to goal start date"))
	}

	return errors
}
