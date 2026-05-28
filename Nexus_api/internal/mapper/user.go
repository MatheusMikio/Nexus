package mapper

import (
	dto "github.com/MatheusMikio/Nexus/internal/domain/dtos/user"
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
)

func UserToResponse(user *schemas.User) *dto.Response {
	return &dto.Response{
		ID:       user.PublicID,
		FullName: user.GetName(),
		Email:    user.GetEmail(),
		Phone:    user.GetPhone(),
	}
}

func UsersToResponse(users []*schemas.User) []*dto.Response {
	usersResponse := make([]*dto.Response, 0, len(users))
	for _, user := range users {
		usersResponse = append(usersResponse, UserToResponse(user))
	}
	return usersResponse
}
