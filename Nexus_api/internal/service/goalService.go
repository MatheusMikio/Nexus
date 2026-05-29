package service

import (
	"errors"

	"github.com/MatheusMikio/Nexus/internal/domain/dtos/goal"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/dates"
	"github.com/MatheusMikio/Nexus/internal/domain/models/parameters"
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"github.com/MatheusMikio/Nexus/internal/mapper"
	"github.com/MatheusMikio/Nexus/internal/repository"
	"github.com/MatheusMikio/Nexus/internal/repository/base"
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
		return nil, userFindError(err)
	}

	goalDb, err := g.GoalRepository.GetByIDAndUserID(id, userDb.ID)
	if err != nil {
		return nil, goalFindError(err)
	}

	return mapper.GoalToResponse(goalDb), nil
}

func (g *GoalService) Create(gr *goal.Request, userId uuid.UUID) []*models.ErrorMessage {
	userDb, err := g.UserRepository.GetByUuid(userId)
	if err != nil {
		return []*models.ErrorMessage{userFindError(err)}
	}

	goalName, errs := models.NewGoalName(gr.Name)
	if len(errs) > 0 {
		return errs
	}

	goalDates, errs := dates.NewGoalDates(gr.StartDate, &gr.EndDate)
	if len(errs) > 0 {
		return errs
	}

	goalDb, errs := schemas.NewGoal(goalName, gr.Description, goalDates, userDb.ID)
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
		return []*models.ErrorMessage{userFindError(err)}
	}

	goalDb, err := g.GoalRepository.GetByIDAndUserID(id, userDb.ID)
	if err != nil {
		return []*models.ErrorMessage{goalFindError(err)}
	}

	if gr.Name != nil {
		goalName, errs := models.NewGoalName(*gr.Name)
		if len(errs) > 0 {
			return errs
		}
		goalDb.GoalName = goalName
	}

	if gr.Description != nil {
		goalDb.Description = *gr.Description
	}

	if gr.StartDate != nil {
		goalDb.Dates.StartDate = *gr.StartDate
	}

	if gr.EndDate != nil {
		goalDb.Dates.FinalizationForecast = gr.EndDate
	}

	if gr.Status != nil {
		status, errMessage := parseGoalStatus(*gr.Status)
		if errMessage != nil {
			return []*models.ErrorMessage{errMessage}
		}
		goalDb.Status = status
	}

	if err := g.GoalRepository.Update(goalDb); err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("Database", err.Error())}
	}

	return nil
}

func (g *GoalService) Delete(id uint, userId uuid.UUID) *models.ErrorMessage {
	userDb, err := g.UserRepository.GetByUuid(userId)
	if err != nil {
		return userFindError(err)
	}

	goalDb, err := g.GoalRepository.GetByIDAndUserID(id, userDb.ID)
	if err != nil {
		return goalFindError(err)
	}

	if err := g.GoalRepository.Delete(goalDb); err != nil {
		return models.NewErrorMessage("Database", err.Error())
	}

	return nil
}

func parseGoalStatus(status string) (schemas.GoalStatus, *models.ErrorMessage) {
	switch schemas.GoalStatus(status) {
	case schemas.GoalPending, schemas.GoalCompleted, schemas.GoalCanceled, schemas.GoalLate:
		return schemas.GoalStatus(status), nil
	default:
		return "", models.NewErrorMessage("Status", "invalid")
	}
}

func goalFindError(err error) *models.ErrorMessage {
	if errors.Is(err, base.ErrNotFound) {
		return models.NewErrorMessage("Goal", "Not found")
	}

	return models.NewErrorMessage("Database", err.Error())
}
