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

func NewUser(fullName models.FullName, contact contact.Contact, password models.Password) (*User, []*models.ErrorMessage) {
	return &User{
		PublicID: uuid.New(),
		FullName: fullName,
		Contact:  contact,
		Password: password,
	}, nil
}
