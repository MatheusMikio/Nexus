package handler

import(
	"net/http"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/service"
	"github.com/gin-gonic/gin"
)
func GetAllUsers(userService service.IUserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params, err := helper.BindPaginationQuery(ctx)
		if err != nil {
			SendError(ctx, http.StatusBadRequest, err)
			return
		}

		users, err := userService.GetAll(params)
		if err != nil {
			SendError(ctx, http.StatusInternalServerError, err)
			return
		}

		SendSuccess(ctx, http.StatusOK, users)
	}
}

func GetUserById(userService service.IUserService) gin.HandlerFunc{
	return func(ctx *gin.Context){
		idStr := ctx.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			SendError(ctx, http.StatusBadRequest, models.NewErrorMessage("Validation", "Invalid user ID"))
			return
		}

		user, err := userService.GetById(id)
		if err != nil{
			SendError(ctx, http.StatusNotFound, err)
			return
		}

		SendSuccess(ctx, http.StatusOK, user)
	}
}