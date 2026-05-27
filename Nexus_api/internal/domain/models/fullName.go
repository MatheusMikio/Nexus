package models

type FullName struct {
	Value string `gorm:"not null;column:name;size:150"`
}
