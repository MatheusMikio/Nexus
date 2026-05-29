package models

type Password struct {
	Value string `gorm:"not null;column:password;size:255"`
}

func NewPassword(value string) (Password, []*ErrorMessage) {
	return Password{
		Value: value,
	}, nil
}

func (p Password) GetValue() string {
	return p.Value
}
