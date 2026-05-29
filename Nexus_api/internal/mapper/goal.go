package mapper

import (
	dto "github.com/MatheusMikio/Nexus/internal/domain/dtos/goal"
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
)

func GoalToResponse(goal *schemas.Goal) *dto.Response {
	tasks := make([]uint, 0, len(goal.Tasks))

	for _, task := range goal.Tasks {
		tasks = append(tasks, task.ID)
	}

	return &dto.Response{
		ID:          goal.ID,
		Name:        goal.GetName(),
		Description: goal.Description,
		StartDate:   goal.GetStartDate(),
		EndDate:     goal.GetFinalDate(),
		Status:      string(goal.Status),
		TaskIDs:     tasks,
	}
}

func GoalsToResponse(goals []*schemas.Goal) []*dto.Response {
	response := make([]*dto.Response, 0, len(goals))

	for _, goal := range goals {
		response = append(response, GoalToResponse(goal))
	}

	return response
}
