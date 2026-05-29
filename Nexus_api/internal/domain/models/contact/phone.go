package contact

import (
	"regexp"
	"strings"

	"github.com/MatheusMikio/Nexus/internal/domain/models"
)

var (
	phoneCleanRegex = regexp.MustCompile(`[()\-\s]`)
	phoneRegex      = regexp.MustCompile(`^\d{10,11}$`)
)

type Phone struct {
	Value string `gorm:"column:phone;not null;size:11"`
}

func NewPhone(value string) (Phone, []*models.ErrorMessage) {
	value = phoneCleanRegex.ReplaceAllString(strings.TrimSpace(value), "")

	if value == "" {
		return Phone{}, []*models.ErrorMessage{
			models.NewErrorMessage("Phone", "is required"),
		}
	}

	if !phoneRegex.MatchString(value) {
		return Phone{}, []*models.ErrorMessage{
			models.NewErrorMessage("Phone", "must have 10 or 11 digits"),
		}
	}

	return Phone{
		Value: value,
	}, nil
}

func (p Phone) GetValue() string {
	return p.Value
}
