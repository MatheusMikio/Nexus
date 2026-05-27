package models

type FullName struct {
	Value string `gorm:"not null;column:name;size:150"`
}

func NewFullName(value string) (FullName, []*ErrorMessage) {
	return FullName{
		Value: value,
	}, nil
}