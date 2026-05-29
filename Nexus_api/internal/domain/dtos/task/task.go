package task

import (
	"time"
)

type Request struct {
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	StartDate        *time.Time `json:"startDate"`
	FinalizationDate *time.Time `json:"finalizationDate"`
}

type Response struct {
	ID               uint       `json:"id"`
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	Status           string     `json:"status"`
	StartDate        *time.Time `json:"startDate"`
	FinalizationDate *time.Time `json:"finalizationDate"`
	TimeSpent        *int64     `json:"timeSpent"`
	GoalID           uint       `json:"goalId"`
}

type Update struct {
	Name             *string    `json:"name"`
	Description      *string    `json:"description"`
	Status           *string    `json:"status"`
	StartDate        *time.Time `json:"startDate"`
	FinalizationDate *time.Time `json:"finalizationDate"`
}
