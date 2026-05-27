package repository

type IGoalRepository interface{
	base.ICrudRepository[schemas.Goal]
	GetAllByUserId(page, size int, userId uuid.UUID) ([]*schemas.Goal, error)
}

type GoalRepository struct{
	base.CrudRepository[schemas.Goal]
}

func NewGoalRepository(db *gorm.DB) IGoalRepository{
	return &GoalRepository{
		CrudRepository: base.CrudRepository[schemas.Goal]
	}
}

func (gr *GoalRepository) GetAllByUserId(page, size int, userId uuid.UUID) ([]*schemas.Goal, error){
	var goals []*schemas.Goal
	offset := (page -1) * size

	if err := gr.Db.Where("user_id = ?", userId).Offset(offset).Limit(size).Find(&goals).Error; err != nil{
		return nil, err
	}
	return goals, nil
}
