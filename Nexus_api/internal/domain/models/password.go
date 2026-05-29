package models

import "unicode"

type Password struct {
	Value string `gorm:"not null;column:password;size:255"`
}

func NewPassword(value string) (Password, []*ErrorMessage) {
	if len(value) < 8 {
		return Password{}, []*ErrorMessage{
			NewErrorMessage("Password", "must have at least 8 characters"),
		}
	}

	hasUppercase := false
	hasNumber := false
	hasSpecial := false

	for _, character := range value {
		switch {
		case unicode.IsUpper(character):
			hasUppercase = true
		case unicode.IsDigit(character):
			hasNumber = true
		case unicode.IsPunct(character) || unicode.IsSymbol(character):
			hasSpecial = true
		}
	}

	errors := make([]*ErrorMessage, 0)

	if !hasUppercase {
		errors = append(errors, NewErrorMessage("Password", "must have at least one uppercase character"))
	}

	if !hasNumber {
		errors = append(errors, NewErrorMessage("Password", "must have at least one number"))
	}

	if !hasSpecial {
		errors = append(errors, NewErrorMessage("Password", "must have at least one special character"))
	}

	if len(errors) > 0 {
		return Password{}, errors
	}

	return Password{
		Value: value,
	}, nil
}

func NewHashedPassword(value string) (Password, []*ErrorMessage) {
	if value == "" {
		return Password{}, []*ErrorMessage{
			NewErrorMessage("Password", "hash is required"),
		}
	}

	if len(value) > 255 {
		return Password{}, []*ErrorMessage{
			NewErrorMessage("Password", "hash must have at most 255 characters"),
		}
	}

	return Password{
		Value: value,
	}, nil
}

func (p Password) GetValue() string {
	return p.Value
}
