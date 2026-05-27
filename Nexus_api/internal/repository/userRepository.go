package repository

import (
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"github.com/MatheusMikio/Nexus/internal/repository/base"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	base.ICrudRepository[schemas.User]
	GetByUuid(uuid uuid.UUID) (*schemas.User, error)
	GetByEmail(email string) (*schemas.User, error)
}

type UserRepository struct {
	base.CrudRepository[schemas.User]
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		CrudRepository: base.CrudRepository[schemas.User]{
			Db: db,
		},
	}
}

func (ur *UserRepository) GetByUuid(uuid uuid.UUID) (*schemas.User, error) {
	var user schemas.User
	if err := ur.Db.Where("public_id = ?", uuid).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetByEmail(email string) (*schemas.User, error) {
	var user schemas.User
	if err := ur.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
