package contact

import (
	"net/mail"
	"strings"

	"github.com/MatheusMikio/Nexus/internal/domain/models"
)

type Email struct {
	Value string `gorm:"column:email;not null;unique;size:255"`
}

func NewEmail(value string) (Email, []*models.ErrorMessage) {
	value = strings.TrimSpace(value)

	if value == "" {
		return Email{}, []*models.ErrorMessage{
			models.NewErrorMessage("Email", "is required"),
		}
	}

	if len(value) > 255 {
		return Email{}, []*models.ErrorMessage{
			models.NewErrorMessage("Email", "must have at most 255 characters"),
		}
	}

	address, err := mail.ParseAddress(value)
	if err != nil || address.Address != value {
		return Email{}, []*models.ErrorMessage{
			models.NewErrorMessage("Email", "invalid"),
		}
	}

	return Email{
		Value: value,
	}, nil
}

func (e Email) GetValue() string {
	return e.Value
}
