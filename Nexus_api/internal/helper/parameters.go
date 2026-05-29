package helper

import (
	"net/http"
	"strconv"

	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/parameters"
	"github.com/MatheusMikio/Nexus/internal/response"
	"github.com/gin-gonic/gin"
)

func BindPaginationQuery(ctx *gin.Context) (parameters.PaginationQuery, *models.ErrorMessage) {
	params := parameters.PaginationQuery{}

	if err := ctx.ShouldBindQuery(&params); err != nil {
		return parameters.PaginationQuery{}, models.NewErrorMessage("Validation", "Invalid query parameters")
	}

	return parameters.NewPaginationQuery(params.Page, params.Size), nil
}

func ParseUintParam(ctx *gin.Context, param string) (uint, bool) {
	id, err := strconv.ParseUint(ctx.Param(param), 10, 64)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, models.NewErrorMessage("ID", "invalid"))
		return 0, false
	}

	return uint(id), true
}
