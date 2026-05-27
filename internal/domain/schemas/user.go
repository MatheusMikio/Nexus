package schemas

import (
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/contact"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName models.FullName `gorm:"embedded"`
	Contact  contact.Contact `gorm:"embedded"`
	Password models.Password `gorm:"embedded"`
	Goals    []Goal
}
