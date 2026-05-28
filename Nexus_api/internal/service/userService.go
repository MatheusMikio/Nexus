package service

import (
	"github.com/MatheusMikio/Nexus/internal/domain/dtos/user"
	"github.com/MatheusMikio/Nexus/internal/domain/factory"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/parameters"
	"github.com/MatheusMikio/Nexus/internal/mapper"
	"github.com/MatheusMikio/Nexus/internal/repository"
	"github.com/google/uuid"
)

type IUserService interface {
	GetAll(parameters parameters.PaginationQuery) ([]*user.Response, *models.ErrorMessage)
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

func (u *UserService) GetAll(parameters parameters.PaginationQuery) ([]*user.Response, *models.ErrorMessage) {
	usersDb, err := u.UserRepo.GetAll(parameters.Page, parameters.Size)

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

func (u *UserService) Create(ur *user.Request) []*models.ErrorMessage {
	user, errors := factory.NewUserFromRequest(ur)

	if len(errors) > 0 {
		return errors
	}

	if err := u.UserRepo.Create(user); err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("Database", err.Error())}
	}

	return nil
}

func (u *UserService) Update(id uuid.UUID, user *user.Update) []*models.ErrorMessage {
	userDb, err := u.UserRepo.GetByUuid(id)

	if err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("User", "Not found")}
	}

	errors := factory.BuildUserUpdate(user, userDb)
	if len(errors) > 0 {
		return errors
	}

	if err := u.UserRepo.Update(userDb); err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("Database", err.Error())}
	}

	return nil
}

func (u *UserService) Delete(id uuid.UUID) *models.ErrorMessage {
	userDb, err := u.UserRepo.GetByUuid(id)

	if err != nil {
		return models.NewErrorMessage("User", "Not found")
	}

	if err := u.UserRepo.Delete(userDb.ID); err != nil {
		return models.NewErrorMessage("Database", err.Error())
	}

	return nil
}
