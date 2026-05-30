package service

import (
	"github.com/MatheusMikio/Nexus/internal/auth"
	"github.com/MatheusMikio/Nexus/internal/domain/dtos/user"
	"github.com/MatheusMikio/Nexus/internal/domain/factory"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/contact"
	"github.com/MatheusMikio/Nexus/internal/domain/models/parameters"
	"github.com/MatheusMikio/Nexus/internal/helper"
	"github.com/MatheusMikio/Nexus/internal/mapper"
	"github.com/MatheusMikio/Nexus/internal/repository"
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
		return nil, helper.FindError("User", err)
	}

	return mapper.UserToResponse(userDb), nil
}

func (u *UserService) CreateUser(ur *user.Request) []*models.ErrorMessage {
	if errors := u.validateNewUserContact(ur.Email, ur.Phone); len(errors) > 0 {
		return errors
	}

	user, errors := factory.NewUserFromRequest(ur)

	if len(errors) > 0 {
		return errors
	}

	hashedPassword, errMessage := hashPassword(user.GetPassword())
	if errMessage != nil {
		return []*models.ErrorMessage{errMessage}
	}

	user.ChangePassword(*hashedPassword)

	if err := u.UserRepo.Create(user); err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("Database", err.Error())}
	}

	return nil
}

func (u *UserService) UpdateUser(id uuid.UUID, user *user.Update) []*models.ErrorMessage {
	userDb, err := u.UserRepo.GetByUuid(id)

	if err != nil {
		return []*models.ErrorMessage{helper.FindError("User", err)}
	}

	if errors := u.validateUpdatedUserContact(user, userDb.GetEmail(), userDb.GetPhone(), userDb.ID); len(errors) > 0 {
		return errors
	}

	errors := factory.BuildUserUpdate(user, userDb)

	if user.Password != nil {
		if _, passwordErrors := models.NewPassword(*user.Password); len(passwordErrors) > 0 {
			errors = append(errors, passwordErrors...)
		}
	}

	if len(errors) > 0 {
		return errors
	}

	if user.Password != nil {
		hashedPassword, errMessage := hashPassword(*user.Password)

		if errMessage != nil {
			return []*models.ErrorMessage{errMessage}
		}

		userDb.ChangePassword(*hashedPassword)
	}

	if err := u.UserRepo.Update(userDb); err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("Database", err.Error())}
	}

	return nil
}

func (u *UserService) DeleteUser(id uuid.UUID) *models.ErrorMessage {
	userDb, err := u.UserRepo.GetByUuid(id)

	if err != nil {
		return helper.FindError("User", err)
	}

	if err := u.UserRepo.Delete(userDb); err != nil {
		return models.NewErrorMessage("Database", err.Error())
	}

	return nil
}

func (u *UserService) validateNewUserContact(email, phone string) []*models.ErrorMessage {
	errors := make([]*models.ErrorMessage, 0)

	newEmail, errMessages := contact.NewEmail(email)
	errors = helper.AppendErrors(errors, errMessages)

	newPhone, errMessages := contact.NewPhone(phone)
	errors = helper.AppendErrors(errors, errMessages)

	if len(errors) > 0 {
		return errors
	}

	emailExists, err := u.UserRepo.ActiveExistsByEmail(newEmail.GetValue())
	if err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("Database", err.Error())}
	}
	if emailExists {
		errors = append(errors, models.NewErrorMessage("Email", "already exists"))
	}

	phoneExists, err := u.UserRepo.ActiveExistsByPhone(newPhone.GetValue())
	if err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("Database", err.Error())}
	}
	if phoneExists {
		errors = append(errors, models.NewErrorMessage("Phone", "already exists"))
	}

	return errors
}

func (u *UserService) validateUpdatedUserContact(user *user.Update, currentEmail, currentPhone string, id uint) []*models.ErrorMessage {
	errors := make([]*models.ErrorMessage, 0)

	if user.Email != nil {
		errors = helper.AppendErrors(errors, u.validateUpdatedEmail(*user.Email, currentEmail, id))
	}

	if user.Phone != nil {
		errors = helper.AppendErrors(errors, u.validateUpdatedPhone(*user.Phone, currentPhone, id))
	}

	return errors
}

func (u *UserService) validateUpdatedEmail(value, currentEmail string, id uint) []*models.ErrorMessage {
	email, errors := contact.NewEmail(value)
	if len(errors) > 0 {
		return errors
	}

	if email.GetValue() == currentEmail {
		return nil
	}

	emailExists, err := u.UserRepo.ActiveExistsByEmailExceptID(email.GetValue(), id)
	if err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("Database", err.Error())}
	}

	if emailExists {
		return []*models.ErrorMessage{models.NewErrorMessage("Email", "already exists")}
	}

	return nil
}

func (u *UserService) validateUpdatedPhone(value, currentPhone string, id uint) []*models.ErrorMessage {
	phone, errors := contact.NewPhone(value)
	if len(errors) > 0 {
		return errors
	}

	if phone.GetValue() == currentPhone {
		return nil
	}

	phoneExists, err := u.UserRepo.ActiveExistsByPhoneExceptID(phone.GetValue(), id)
	if err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("Database", err.Error())}
	}

	if phoneExists {
		return []*models.ErrorMessage{models.NewErrorMessage("Phone", "already exists")}
	}

	return nil
}

func hashPassword(rawPassword string) (*models.Password, *models.ErrorMessage) {
	passwordHash, err := auth.HashPassword(rawPassword)
	if err != nil {
		return nil, models.NewErrorMessage("Password", "failed to hash")
	}

	hashedPassword, errors := models.NewHashedPassword(passwordHash)
	if len(errors) > 0 {
		return nil, errors[0]
	}

	return &hashedPassword, nil
}
