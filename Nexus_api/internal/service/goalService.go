package service

import (
	"github.com/MatheusMikio/Nexus/internal/domain/dtos/goal"
	"github.com/MatheusMikio/Nexus/internal/domain/factory"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/parameters"
	"github.com/MatheusMikio/Nexus/internal/helper"
	"github.com/MatheusMikio/Nexus/internal/mapper"
	"github.com/MatheusMikio/Nexus/internal/repository"
	"github.com/google/uuid"
)

type IGoalService interface {
	GetAllGoals(parameters parameters.PaginationQuery, userId uuid.UUID) ([]*goal.Response, *models.ErrorMessage)
	GetById(id uint, userId uuid.UUID) (*goal.Response, *models.ErrorMessage)
	Create(goal *goal.Request, userId uuid.UUID) []*models.ErrorMessage
	Update(id uint, goal *goal.Update, userId uuid.UUID) []*models.ErrorMessage
	Delete(id uint, userId uuid.UUID) *models.ErrorMessage
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

func (g *GoalService) GetAllGoals(parameters parameters.PaginationQuery, userId uuid.UUID) ([]*goal.Response, *models.ErrorMessage) {
	goalsDb, err := g.GoalRepository.GetAllByUserId(parameters.Page, parameters.Size, userId)
	if err != nil {
		return nil, models.NewErrorMessage("Database", err.Error())
	}

	return mapper.GoalsToResponse(goalsDb), nil
}

func (g *GoalService) GetById(id uint, userId uuid.UUID) (*goal.Response, *models.ErrorMessage) {
	userDb, err := g.UserRepository.GetByUuid(userId)
	if err != nil {
		return nil, helper.FindError("User", err)
	}

	goalDb, err := g.GoalRepository.GetByIDAndUserID(id, userDb.ID)
	if err != nil {
		return nil, helper.FindError("Goal", err)
	}

	return mapper.GoalToResponse(goalDb), nil
}

func (g *GoalService) Create(gr *goal.Request, userId uuid.UUID) []*models.ErrorMessage {
	userDb, err := g.UserRepository.GetByUuid(userId)
	if err != nil {
		return []*models.ErrorMessage{helper.FindError("User", err)}
	}

	goalDb, errs := factory.NewGoalFromRequest(gr, userDb.ID)
	if len(errs) > 0 {
		return errs
	}

	if err := g.GoalRepository.Create(goalDb); err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("Database", err.Error())}
	}

	return nil
}

func (g *GoalService) Update(id uint, gr *goal.Update, userId uuid.UUID) []*models.ErrorMessage {
	userDb, err := g.UserRepository.GetByUuid(userId)
	if err != nil {
		return []*models.ErrorMessage{helper.FindError("User", err)}
	}

	goalDb, err := g.GoalRepository.GetByIDAndUserID(id, userDb.ID)
	if err != nil {
		return []*models.ErrorMessage{helper.FindError("Goal", err)}
	}

	errs := factory.BuildGoalUpdate(gr, goalDb)
	if len(errs) > 0 {
		return errs
	}

	if err := g.GoalRepository.Update(goalDb); err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("Database", err.Error())}
	}

	return nil
}

func (g *GoalService) Delete(id uint, userId uuid.UUID) *models.ErrorMessage {
	userDb, err := g.UserRepository.GetByUuid(userId)
	if err != nil {
		return helper.FindError("User", err)
	}

	goalDb, err := g.GoalRepository.GetByIDAndUserID(id, userDb.ID)
	if err != nil {
		return helper.FindError("Goal", err)
	}

	if err := g.GoalRepository.Delete(goalDb); err != nil {
		return models.NewErrorMessage("Database", err.Error())
	}

	return nil
}
