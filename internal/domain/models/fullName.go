package models

type FullName struct {
	Value string `gorm:"not null;size:150" json:"fullName"`
}
