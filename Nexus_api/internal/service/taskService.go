package service

import (
	"github.com/MatheusMikio/Nexus/internal/domain/dtos/task"
	"github.com/MatheusMikio/Nexus/internal/domain/factory"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/parameters"
	"github.com/MatheusMikio/Nexus/internal/helper"
	"github.com/MatheusMikio/Nexus/internal/mapper"
	"github.com/MatheusMikio/Nexus/internal/repository"
	"github.com/google/uuid"
)

type ITaskService interface {
	GetAllTasks(parameters parameters.PaginationQuery, goalId uint, userId uuid.UUID) ([]*task.Response, *models.ErrorMessage)
	GetById(goalId, id uint, userId uuid.UUID) (*task.Response, *models.ErrorMessage)
	Create(goalID uint, task *task.Request, userId uuid.UUID) []*models.ErrorMessage
	Update(goalId, id uint, task *task.Update, userId uuid.UUID) []*models.ErrorMessage
	Delete(goalId, id uint, userId uuid.UUID) *models.ErrorMessage
}

type TaskService struct {
	TaskRepository repository.ITaskRepository
	GoalRepository repository.IGoalRepository
	UserRepository repository.IUserRepository
}

func NewTaskService(taskRepo repository.ITaskRepository, goalRepo repository.IGoalRepository, userRepo repository.IUserRepository) ITaskService {
	return &TaskService{
		TaskRepository: taskRepo,
		GoalRepository: goalRepo,
		UserRepository: userRepo,
	}
}

func (t *TaskService) GetAllTasks(parameters parameters.PaginationQuery, goalId uint, userId uuid.UUID) ([]*task.Response, *models.ErrorMessage) {
	user, err := t.UserRepository.GetByUuid(userId)
	if err != nil {
		return nil, helper.FindError("User", err)
	}

	tasksDb, err := t.TaskRepository.GetAllByGoalID(
		parameters.Page,
		parameters.Size,
		goalId,
		user.ID,
	)

	if err != nil {
		return nil, models.NewErrorMessage("Database", err.Error())
	}

	return mapper.TasksToResponse(tasksDb), nil
}

func (t *TaskService) GetById(goalId, id uint, userId uuid.UUID) (*task.Response, *models.ErrorMessage) {
	userDb, err := t.UserRepository.GetByUuid(userId)
	if err != nil {
		return nil, helper.FindError("User", err)
	}

	taskDb, err := t.TaskRepository.GetByIDAndGoalIDAndUserID(id, goalId, userDb.ID)
	if err != nil {
		return nil, helper.FindError("Task", err)
	}

	return mapper.TaskToResponse(taskDb), nil
}

func (t *TaskService) Create(goalID uint, tr *task.Request, userId uuid.UUID) []*models.ErrorMessage {
	userDb, err := t.UserRepository.GetByUuid(userId)
	if err != nil {
		return []*models.ErrorMessage{helper.FindError("User", err)}
	}

	goalDb, err := t.GoalRepository.GetByIDAndUserID(goalID, userDb.ID)
	if err != nil {
		return []*models.ErrorMessage{helper.FindError("Goal", err)}
	}

	taskDb, errs := factory.NewTaskFromRequest(tr, goalDb)
	if len(errs) > 0 {
		return errs
	}

	if err := t.TaskRepository.Create(taskDb); err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("Database", err.Error())}
	}

	return nil
}

func (t *TaskService) Update(goalId, id uint, tr *task.Update, userId uuid.UUID) []*models.ErrorMessage {
	userDb, err := t.UserRepository.GetByUuid(userId)
	if err != nil {
		return []*models.ErrorMessage{helper.FindError("User", err)}
	}

	goalDb, err := t.GoalRepository.GetByIDAndUserID(goalId, userDb.ID)
	if err != nil {
		return []*models.ErrorMessage{helper.FindError("Goal", err)}
	}

	taskDb, err := t.TaskRepository.GetByIDAndGoalIDAndUserID(id, goalId, userDb.ID)
	if err != nil {
		return []*models.ErrorMessage{helper.FindError("Task", err)}
	}

	errs := factory.BuildTaskUpdate(tr, taskDb, goalDb)
	if len(errs) > 0 {
		return errs
	}

	if err := t.TaskRepository.Update(taskDb); err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("Database", err.Error())}
	}

	return nil
}

func (t *TaskService) Delete(goalId, id uint, userId uuid.UUID) *models.ErrorMessage {
	userDb, err := t.UserRepository.GetByUuid(userId)
	if err != nil {
		return helper.FindError("User", err)
	}

	taskDb, err := t.TaskRepository.GetByIDAndGoalIDAndUserID(id, goalId, userDb.ID)
	if err != nil {
		return helper.FindError("Task", err)
	}

	if err := t.TaskRepository.Delete(taskDb); err != nil {
		return models.NewErrorMessage("Database", err.Error())
	}

	return nil
}
