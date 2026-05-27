package models

type Password struct {
	Value string `gorm:"not null;column:password;size:8"`
}

func NewPassword(value string) (Password, []*ErrorMessage){
	return Password{
		Value: value,
	}, nil
}