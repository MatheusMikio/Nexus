package models

type GoalName struct {
	Value string `gorm:"not null;column:name;size:150"`
}

func NewGoalName(value string) (GoalName, []*ErrorMessage) {
	return GoalName{
		Value: value,
	}, nil
}