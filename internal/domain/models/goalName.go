package models

type GoalName struct {
	Value string `gorm:"not null;column:name;size:150"`
}
