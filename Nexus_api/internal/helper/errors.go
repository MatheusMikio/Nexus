package helper

import (
	"net/http"

	"github.com/MatheusMikio/Nexus/internal/domain/models"
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
