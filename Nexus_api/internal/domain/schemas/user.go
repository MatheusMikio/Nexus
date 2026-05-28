package schemas

import (
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/contact"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PublicID uuid.UUID       `gorm:"type:uuid;unique;not null"`
	FullName models.FullName `gorm:"embedded"`
	Contact  contact.Contact `gorm:"embedded"`
	Password models.Password `gorm:"embedded"`
	Goals    []Goal
}

func NewUser(fullName models.FullName, contact contact.Contact, password models.Password) *User {
	return &User{
		PublicID: uuid.New(),
		FullName: fullName,
		Contact:  contact,
		Password: password,
	}
}

func (u *User) GetName() string {
	return u.FullName.GetValue()
}

func (u *User) GetEmail() string {
	return u.Contact.GetEmail()
}

func (u *User) GetPhone() string {
	return u.Contact.GetPhone()
}

func (u *User) GetPassword() string {
	return u.Password.GetValue()
}


func (u *User) Update(fullName *models.FullName, contact *contact.Contact, password *models.Password) {
	if fullName != nil {
		u.FullName = *fullName
	}

	if contact != nil {
		u.Contact = *contact
	}

	if password != nil {
		u.Password = *password
	}
}
