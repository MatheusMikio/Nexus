package service

import (
	"github.com/MatheusMikio/Nexus/internal/domain/dtos/user"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/mapper"
	"github.com/MatheusMikio/Nexus/internal/repository"
	"github.com/google/uuid"
)

type IUserService interface {
	GetAll(page, size int) ([]*user.Response, *models.ErrorMessage)
	GetById(id uuid.UUID) (*user.Response, *models.ErrorMessage)
	Create(user *user.Request) []*models.ErrorMessage
	Update(id uuid.UUID, user *user.Update) []*models.ErrorMessage
	Delete(id uuid.UUID) *models.ErrorMessage
}

type UserService struct {
	UserRepo repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (u *UserService) GetAll(page int, size int) ([]*user.Response, *models.ErrorMessage) {
	usersDb, err := u.UserRepo.GetAll(page, size)
	if err != nil {
		return nil, models.NewErrorMessage("Database", err.Error())
	}
	return mapper.UsersToResponse(usersDb), nil
}

func (u *UserService) GetById(id uuid.UUID) (*user.Response, *models.ErrorMessage) {
	userDb, err := u.UserRepo.GetByUuid(id)
	if err != nil {
		return nil, models.NewErrorMessage("User", "Not found")
	}
	return mapper.UserToResponse(userDb), nil
}

func (u *UserService) Create(user *user.Request) []*models.ErrorMessage {
	panic("unimplemented")
}

func (u *UserService) Update(id uuid.UUID, user *user.Update) []*models.ErrorMessage {
	panic("unimplemented")
}

func (u *UserService) Delete(id uuid.UUID) *models.ErrorMessage {
	panic("unimplemented")
}
