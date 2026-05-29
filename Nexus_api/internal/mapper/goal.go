package mapper

import (
	dto "github.com/MatheusMikio/Nexus/internal/domain/dtos/goal"
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
)

func GoalToResponse(goal *schemas.Goal) *dto.Response {
	return &dto.Response{
		ID:          goal.ID,
		Name:        goal.GetName(),
		Description: goal.GetDescription(),
		StartDate:   goal.GetStartDate(),
		EndDate:     goal.GetFinalDate(),
		Status:      string(goal.GetStatus()),
		TaskIDs:     goal.GetTaskIDs(),
	}
}

func GoalsToResponse(goals []*schemas.Goal) []*dto.Response {
	response := make([]*dto.Response, 0, len(goals))

	for _, goal := range goals {
		response = append(response, GoalToResponse(goal))
	}

	return response
}
