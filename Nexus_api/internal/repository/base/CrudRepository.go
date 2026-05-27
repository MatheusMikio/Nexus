package base

import "gorm.io/gorm"

type ICrudRepository[T any] interface {
	GetAll(page, size int) ([]*T, error)
	GetByID(id uint) (*T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
}

type CrudRepository[T any] struct {
	Db *gorm.DB
}

func (cr *CrudRepository[T]) GetAll(page, size int) ([]*T, error) {
	var entities []*T
	offset := (page - 1) * size

	if err := cr.Db.Offset(offset).Limit(size).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (cr *CrudRepository[T]) GetByID(id uint) (*T, error) {
	var entity T
	if err := cr.Db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (cr *CrudRepository[T]) Create(entity *T) error {
	if err := cr.Db.Create(entity).Error; err != nil {
		return err
	}

	return nil
}

func (cr *CrudRepository[T]) Update(entity *T) error {
	if err := cr.Db.Save(entity).Error; err != nil {
		return err
	}

	return nil
}

func (cr *CrudRepository[T]) Delete(entity *T) error {
	if err := cr.Db.Delete(entity).Error; err != nil {
		return err
	}

	return nil
}
