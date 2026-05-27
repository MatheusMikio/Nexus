package models

type Password struct {
	Value string `gorm:"not null;size:8" json:"password"`
}
