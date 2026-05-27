package service

import(
	"github.com/google/uuid"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/repository"
	"github.com/MatheusMikio/Nexus/internal/domain/dtos/user"
)

type IUserService interface{
	GetAll(page, size int) ([]*user.Response, *models.ErrorMessage)
	GetById(id uuid.UUID) (*user.Response, *models.ErrorMessage)
	Create(user *user.Request) []*models.ErrorMessage
	Update(id uuid.UUID, user *user.Update) []*models.ErrorMessage
	Delete(id uuid.UUID) *models.ErrorMessage
}

type UserService struct{
	UserRepo repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) IUserService{
	return &UserService{
		UserRepo: userRepo
	}
}