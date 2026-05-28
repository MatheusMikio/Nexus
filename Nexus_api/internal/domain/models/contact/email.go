package contact

import "github.com/MatheusMikio/Nexus/internal/domain/models"

type Email struct {
	Value string `gorm:"column:email;not null;unique;size:255"`
}

func NewEmail(value string) (Email, []*models.ErrorMessage) {
	return Email{
		Value: value,
	}, nil
}

func (e Email) GetValue() string {
	return e.Value
}
