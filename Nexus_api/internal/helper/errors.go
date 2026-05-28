package helper

import (
	"github.com/MatheusMikio/Nexus/internal/domain/models"
)

func AppendErrors(target []*models.ErrorMessage, newErrors []*models.ErrorMessage) []*models.ErrorMessage {
	if len(newErrors) > 0 {
		return append(target, newErrors...)
	}
	return target
}
