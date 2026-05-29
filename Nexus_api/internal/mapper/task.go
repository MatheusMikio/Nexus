package mapper

import (
	dto "github.com/MatheusMikio/Nexus/internal/domain/dtos/task"
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
)

func TaskToResponse(task *schemas.Task) *dto.Response {
	return &dto.Response{
		ID:               task.ID,
		Name:             task.GetName(),
		Description:      task.Description,
		Status:           string(task.Status),
		StartDate:        task.GetStartDate(),
		FinalizationDate: task.GetFinalizationDate(),
		TimeSpent:        task.GetTimeSpent(),
		GoalID:           task.GoalID,
	}
}

func TasksToResponse(tasks []*schemas.Task) []*dto.Response {
	response := make([]*dto.Response, 0, len(tasks))
	for _, task := range tasks {
		response = append(response, TaskToResponse(task))
	}
	return response
}
