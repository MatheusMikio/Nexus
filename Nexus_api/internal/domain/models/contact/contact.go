package contact

import "github.com/MatheusMikio/Nexus/internal/domain/models"

type Contact struct {
	Email Email `gorm:"embedded"`
	Phone Phone `gorm:"embedded"`
}

func NewContact(email Email, phone Phone) (Contact, []*models.ErrorMessage) {
	return Contact{
		Email: email,
		Phone: phone,
	}, nil
}

func (c Contact) GetEmail() string {
	return c.Email.GetValue()
}

func (c Contact) GetPhone() string {
	return c.Phone.GetValue()
}
