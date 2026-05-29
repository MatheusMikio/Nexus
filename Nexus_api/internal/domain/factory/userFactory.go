package factory

import (
	dto "github.com/MatheusMikio/Nexus/internal/domain/dtos/user"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/contact"
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"github.com/MatheusMikio/Nexus/internal/helper"
)

func NewUserFromRequest(ur *dto.Request) (*schemas.User, []*models.ErrorMessage) {
	errors := make([]*models.ErrorMessage, 0)

	fullname, err := models.NewFullName(ur.FullName)
	errors = helper.AppendErrors(errors, err)

	email, err := contact.NewEmail(ur.Email)
	errors = helper.AppendErrors(errors, err)

	phone, err := contact.NewPhone(ur.Phone)
	errors = helper.AppendErrors(errors, err)

	contact, err := contact.NewContact(email, phone)
	errors = helper.AppendErrors(errors, err)

	password, err := models.NewPassword(ur.Password)
	errors = helper.AppendErrors(errors, err)

	if len(errors) > 0 {
		return nil, errors
	}

	return schemas.NewUser(fullname, contact, password), nil
}

func BuildUserUpdate(user *dto.Update, userDb *schemas.User) []*models.ErrorMessage {
	errors := make([]*models.ErrorMessage, 0)

	name, err := buildFullName(user.FullName, userDb.FullName)
	errors = helper.AppendErrors(errors, err)

	contact, err := buildContact(user, userDb.Contact)
	errors = helper.AppendErrors(errors, err)

	if len(errors) > 0 {
		return errors
	}

	userDb.Update(name, contact, nil)
	return nil
}

func buildFullName(value *string, current models.FullName) (*models.FullName, []*models.ErrorMessage) {
	if value == nil || *value == current.GetValue() {
		return nil, nil
	}

	fullName, errors := models.NewFullName(*value)
	if len(errors) > 0 {
		return nil, errors
	}

	return &fullName, nil
}

func buildContact(user *dto.Update, current contact.Contact) (*contact.Contact, []*models.ErrorMessage) {
	emailValue := current.GetEmail()
	phoneValue := current.GetPhone()

	if user.Email != nil {
		emailValue = *user.Email
	}

	if user.Phone != nil {
		phoneValue = *user.Phone
	}

	if emailValue == current.GetEmail() && phoneValue == current.GetPhone() {
		return nil, nil
	}

	errors := make([]*models.ErrorMessage, 0)

	email, err := contact.NewEmail(emailValue)
	errors = helper.AppendErrors(errors, err)

	phone, err := contact.NewPhone(phoneValue)
	errors = helper.AppendErrors(errors, err)

	newContact, err := contact.NewContact(email, phone)
	errors = helper.AppendErrors(errors, err)

	if len(errors) > 0 {
		return nil, errors
	}

	return &newContact, nil
}
