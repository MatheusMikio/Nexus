package models

type Password struct {
	Value string `gorm:"not null;column:password;size:8"`
}
