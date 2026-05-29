package models

import (
	"strings"
	"unicode/utf8"
)

type GoalName struct {
	Value string `gorm:"not null;column:name;size:150"`
}

func NewGoalName(value string) (GoalName, []*ErrorMessage) {
	value = strings.TrimSpace(value)

	length := utf8.RuneCountInString(value)

	if length <= 5 {
		return GoalName{}, []*ErrorMessage{
			NewErrorMessage("Name", "must have more than 5 characters"),
		}
	}

	if length > 150 {
		return GoalName{}, []*ErrorMessage{
			NewErrorMessage("Name", "must have at most 150 characters"),
		}
	}

	return GoalName{
		Value: value,
	}, nil
}

func (g *GoalName) GetValue() string {
	return g.Value
}
