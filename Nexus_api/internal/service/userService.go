package service

import (
	"errors"

	"github.com/MatheusMikio/Nexus/internal/domain/dtos/user"
	"github.com/MatheusMikio/Nexus/internal/domain/factory"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/parameters"
	"github.com/MatheusMikio/Nexus/internal/mapper"
	"github.com/MatheusMikio/Nexus/internal/repository"
	"github.com/MatheusMikio/Nexus/internal/repository/base"
	"github.com/google/uuid"
)

type IUserService interface {
	GetAllUsers(parameters parameters.PaginationQuery) ([]*user.Response, *models.ErrorMessage)
	GetUserById(id uuid.UUID) (*user.Response, *models.ErrorMessage)
	CreateUser(user *user.Request) []*models.ErrorMessage
	UpdateUser(id uuid.UUID, user *user.Update) []*models.ErrorMessage
	DeleteUser(id uuid.UUID) *models.ErrorMessage
}

type UserService struct {
	UserRepo repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (u *UserService) GetAllUsers(parameters parameters.PaginationQuery) ([]*user.Response, *models.ErrorMessage) {
	usersDb, err := u.UserRepo.GetAllWithGoals(parameters.Page, parameters.Size)

	if err != nil {
		return nil, models.NewErrorMessage("Database", err.Error())
	}

	return mapper.UsersToResponse(usersDb), nil
}

func (u *UserService) GetUserById(id uuid.UUID) (*user.Response, *models.ErrorMessage) {
	userDb, err := u.UserRepo.GetByUuidWithGoals(id)

	if err != nil {
		return nil, userFindError(err)
	}

	return mapper.UserToResponse(userDb), nil
}

func (u *UserService) CreateUser(ur *user.Request) []*models.ErrorMessage {
	user, errors := factory.NewUserFromRequest(ur)

	if len(errors) > 0 {
		return errors
	}

	if err := u.UserRepo.Create(user); err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("Database", err.Error())}
	}

	return nil
}

func (u *UserService) UpdateUser(id uuid.UUID, user *user.Update) []*models.ErrorMessage {
	userDb, err := u.UserRepo.GetByUuid(id)

	if err != nil {
		return []*models.ErrorMessage{userFindError(err)}
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

func (u *UserService) DeleteUser(id uuid.UUID) *models.ErrorMessage {
	userDb, err := u.UserRepo.GetByUuid(id)

	if err != nil {
		return userFindError(err)
	}

	if err := u.UserRepo.Delete(userDb); err != nil {
		return models.NewErrorMessage("Database", err.Error())
	}

	return nil
}

func userFindError(err error) *models.ErrorMessage {
	if errors.Is(err, base.ErrNotFound) {
		return models.NewErrorMessage("User", "Not found")
	}

	return models.NewErrorMessage("Database", err.Error())
}
