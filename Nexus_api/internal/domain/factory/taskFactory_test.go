package factory

import (
	"testing"
	"time"

	dto "github.com/MatheusMikio/Nexus/internal/domain/dtos/task"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/dates"
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
)

func TestBuildTaskUpdateAllowsCompletedTaskWithUnchangedPastStartDate(t *testing.T) {
	now := time.Now().UTC()
	startDate := now.AddDate(0, 0, -1)
	finalizationDate := now
	status := string(schemas.TaskCompleted)

	taskDb := buildTaskForUpdateTest(t, startDate)
	goal := buildGoalForTaskUpdateTest(t, startDate.AddDate(0, 0, -1))

	errs := BuildTaskUpdate(&dto.Update{
		Status:           &status,
		StartDate:        &startDate,
		FinalizationDate: &finalizationDate,
	}, taskDb, goal)

	if len(errs) > 0 {
		t.Fatalf("expected unchanged past start date to be accepted, got errors: %#v", errs)
	}

	if taskDb.GetStatus() != schemas.TaskCompleted {
		t.Fatalf("expected task status %q, got %q", schemas.TaskCompleted, taskDb.GetStatus())
	}
}

func TestBuildTaskUpdateRejectsChangedPastStartDate(t *testing.T) {
	now := time.Now().UTC()
	currentStartDate := now
	pastStartDate := now.AddDate(0, 0, -1)

	taskDb := buildTaskForUpdateTest(t, currentStartDate)
	goal := buildGoalForTaskUpdateTest(t, pastStartDate.AddDate(0, 0, -1))

	errs := BuildTaskUpdate(&dto.Update{
		StartDate: &pastStartDate,
	}, taskDb, goal)

	if !hasErrorMessage(errs, "StartDate", "must be greater than or equal to current date") {
		t.Fatalf("expected changed past start date to be rejected, got errors: %#v", errs)
	}
}

func buildTaskForUpdateTest(t *testing.T, startDate time.Time) *schemas.Task {
	t.Helper()

	taskName, errs := models.NewGoalName("Existing task")
	if len(errs) > 0 {
		t.Fatalf("expected valid task name, got errors: %#v", errs)
	}

	taskDates, errs := dates.NewTaskDates(&startDate, nil)
	if len(errs) > 0 {
		t.Fatalf("expected valid task dates, got errors: %#v", errs)
	}

	return &schemas.Task{
		Name:        taskName,
		Description: "Existing description",
		Status:      schemas.TaskPending,
		Dates:       taskDates,
		GoalID:      1,
	}
}

func buildGoalForTaskUpdateTest(t *testing.T, startDate time.Time) *schemas.Goal {
	t.Helper()

	goalName, errs := models.NewGoalName("Existing goal")
	if len(errs) > 0 {
		t.Fatalf("expected valid goal name, got errors: %#v", errs)
	}

	return &schemas.Goal{
		GoalName: goalName,
		Dates: dates.GoalDates{
			StartDate:            startDate,
			FinalizationForecast: startDate.AddDate(0, 0, 30),
		},
		Status: schemas.GoalPending,
		UserID: 1,
	}
}

func hasErrorMessage(errs []*models.ErrorMessage, field string, message string) bool {
	for _, err := range errs {
		if err.Property == field && err.Message == message {
			return true
		}
	}

	return false
}
