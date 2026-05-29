package models

import (
	"strings"
	"unicode/utf8"
)

type FullName struct {
	Value string `gorm:"not null;column:name;size:150"`
}

func NewFullName(value string) (FullName, []*ErrorMessage) {
	value = strings.TrimSpace(value)

	length := utf8.RuneCountInString(value)

	if length < 10 {
		return FullName{}, []*ErrorMessage{
			NewErrorMessage("FullName", "must have at least 10 characters"),
		}
	}

	if length > 150 {
		return FullName{}, []*ErrorMessage{
			NewErrorMessage("FullName", "must have at most 150 characters"),
		}
	}

	return FullName{
		Value: value,
	}, nil
}

func (f FullName) GetValue() string {
	return f.Value
}
