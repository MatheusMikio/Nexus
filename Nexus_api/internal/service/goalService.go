package service

import (
	"github.com/MatheusMikio/Nexus/internal/domain/dtos/goal"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/repository"
	"github.com/google/uuid"
)

type IGoalService interface {
	GetAllGoals(page, size int, userID uuid.UUID) ([]*goal.Response, *models.ErrorMessage)
	GetById(id uint, userID uuid.UUID) (*goal.Response, *models.ErrorMessage)
	Create(goal *goal.Request, userID uuid.UUID) []*models.ErrorMessage
	Update(id uint, goal *goal.Update, userID uuid.UUID) []*models.ErrorMessage
	Delete(id uint, userID uuid.UUID) *models.ErrorMessage
}

type GoalService struct {
	GoalRepository repository.IGoalRepository
	UserRepository repository.IUserRepository
}

func NewGoalService(goalRepo repository.IGoalRepository, userRepo repository.IUserRepository) IGoalService {
	return &GoalService{
		GoalRepository: goalRepo,
		UserRepository: userRepo,
	}
}

func (g *GoalService) GetAllGoals(page int, size int, userID uuid.UUID) ([]*goal.Response, *models.ErrorMessage) {
	panic("unimplemented")
}

func (g *GoalService) GetById(id uint, userID uuid.UUID) (*goal.Response, *models.ErrorMessage) {
	panic("unimplemented")
}

func (g *GoalService) Create(goal *goal.Request, userID uuid.UUID) []*models.ErrorMessage {
	panic("unimplemented")
}

func (g *GoalService) Update(id uint, goal *goal.Update, userID uuid.UUID) []*models.ErrorMessage {
	panic("unimplemented")
}

func (g *GoalService) Delete(id uint, userID uuid.UUID) *models.ErrorMessage {
	panic("unimplemented")
}
