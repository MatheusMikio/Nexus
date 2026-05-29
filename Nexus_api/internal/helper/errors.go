package helper

import (
	"errors"
	"net/http"

	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/repository/base"
)

func AppendErrors(target []*models.ErrorMessage, newErrors []*models.ErrorMessage) []*models.ErrorMessage {
	if len(newErrors) > 0 {
		return append(target, newErrors...)
	}
	return target
}

func ErrorStatusCode(err *models.ErrorMessage) int {
	if err.Message == "Not found" {
		return http.StatusNotFound
	}

	if err.Property == "Database" {
		return http.StatusInternalServerError
	}

	return http.StatusBadRequest
}

func FindError(entity string, err error) *models.ErrorMessage {
	if errors.Is(err, base.ErrNotFound) {
		return models.NewErrorMessage(entity, "Not found")
	}

	return models.NewErrorMessage("Database", err.Error())
}
