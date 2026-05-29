package repository

import (
	"errors"

	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"github.com/MatheusMikio/Nexus/internal/repository/base"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	base.ICrudRepository[schemas.User]
	GetAllWithGoals(page, size int) ([]*schemas.User, error)
	GetByUuid(uuid uuid.UUID) (*schemas.User, error)
	GetByUuidWithGoals(uuid uuid.UUID) (*schemas.User, error)
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, base.ErrNotFound
		}

		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetAllWithGoals(page, size int) ([]*schemas.User, error) {
	var users []*schemas.User
	offset := (page - 1) * size

	if err := ur.Db.
		Preload("Goals").
		Offset(offset).
		Limit(size).
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) GetByUuidWithGoals(uuid uuid.UUID) (*schemas.User, error) {
	var user schemas.User
	if err := ur.Db.Preload("Goals").Where("public_id = ?", uuid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, base.ErrNotFound
		}

		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetByEmail(email string) (*schemas.User, error) {
	var user schemas.User
	if err := ur.Db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, base.ErrNotFound
		}

		return nil, err
	}
	return &user, nil
}
