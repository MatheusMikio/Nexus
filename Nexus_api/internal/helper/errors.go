package helper

import (
	"errors"
	"net/http"

	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/repository/base"
	"github.com/jackc/pgx/v5/pgconn"
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

func IsUniqueViolation(err error, constraintNames ...string) bool {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) || pgErr.Code != "23505" {
		return false
	}

	if len(constraintNames) == 0 {
		return true
	}

	for _, constraintName := range constraintNames {
		if pgErr.ConstraintName == constraintName {
			return true
		}
	}

	return false
}
