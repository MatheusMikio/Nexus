package helper

import (
	"github.com/MatheusMikio/Nexus/internal/domain/dtos/parameters"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/gin-gonic/gin"
)

func BindPaginationQuery(ctx *gin.Context) (parameters.PaginationQuery, *models.ErrorMessage) {
	params := parameters.PaginationQuery{}

	if err := ctx.ShouldBindQuery(&params); err != nil {
		return parameters.PaginationQuery{}, models.NewErrorMessage("Validation", "Invalid query parameters")
	}

	return parameters.NewPaginationQuery(params.Page, params.Size), nil
}