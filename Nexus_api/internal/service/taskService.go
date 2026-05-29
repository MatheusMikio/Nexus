package service

import (
	"errors"

	"github.com/MatheusMikio/Nexus/internal/domain/dtos/task"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/dates"
	"github.com/MatheusMikio/Nexus/internal/domain/models/parameters"
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"github.com/MatheusMikio/Nexus/internal/mapper"
	"github.com/MatheusMikio/Nexus/internal/repository"
	"github.com/MatheusMikio/Nexus/internal/repository/base"
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
		return nil, models.NewErrorMessage("User", "Not found.")
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
		return nil, userFindError(err)
	}

	taskDb, err := t.TaskRepository.GetByIDAndGoalIDAndUserID(id, goalId, userDb.ID)
	if err != nil {
		return nil, taskFindError(err)
	}

	return mapper.TaskToResponse(taskDb), nil
}

func (t *TaskService) Create(goalID uint, tr *task.Request, userId uuid.UUID) []*models.ErrorMessage {
	userDb, err := t.UserRepository.GetByUuid(userId)
	if err != nil {
		return []*models.ErrorMessage{userFindError(err)}
	}

	if _, err := t.GoalRepository.GetByIDAndUserID(goalID, userDb.ID); err != nil {
		return []*models.ErrorMessage{goalFindError(err)}
	}

	taskName, errs := models.NewGoalName(tr.Name)
	if len(errs) > 0 {
		return errs
	}

	taskDates, errs := dates.NewTaskDates(tr.StartDate, tr.FinalizationDate)
	if len(errs) > 0 {
		return errs
	}

	taskDb, errs := schemas.NewTask(taskName, tr.Description, taskDates, goalID)
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
		return []*models.ErrorMessage{userFindError(err)}
	}

	taskDb, err := t.TaskRepository.GetByIDAndGoalIDAndUserID(id, goalId, userDb.ID)
	if err != nil {
		return []*models.ErrorMessage{taskFindError(err)}
	}

	if tr.Name != nil {
		taskName, errs := models.NewGoalName(*tr.Name)
		if len(errs) > 0 {
			return errs
		}
		taskDb.Name = taskName
	}

	if tr.Description != nil {
		taskDb.Description = *tr.Description
	}

	if tr.Status != nil {
		status, errMessage := parseTaskStatus(*tr.Status)
		if errMessage != nil {
			return []*models.ErrorMessage{errMessage}
		}
		taskDb.Status = status
	}

	startDate := taskDb.Dates.StartDate
	if tr.StartDate != nil {
		startDate = tr.StartDate
	}

	finalizationDate := taskDb.Dates.FinalizationDate
	if tr.FinalizationDate != nil {
		finalizationDate = tr.FinalizationDate
	}

	taskDates, errs := dates.NewTaskDates(startDate, finalizationDate)
	if len(errs) > 0 {
		return errs
	}
	taskDb.Dates = taskDates

	if err := t.TaskRepository.Update(taskDb); err != nil {
		return []*models.ErrorMessage{models.NewErrorMessage("Database", err.Error())}
	}

	return nil
}

func (t *TaskService) Delete(goalId, id uint, userId uuid.UUID) *models.ErrorMessage {
	userDb, err := t.UserRepository.GetByUuid(userId)
	if err != nil {
		return userFindError(err)
	}

	taskDb, err := t.TaskRepository.GetByIDAndGoalIDAndUserID(id, goalId, userDb.ID)
	if err != nil {
		return taskFindError(err)
	}

	if err := t.TaskRepository.Delete(taskDb); err != nil {
		return models.NewErrorMessage("Database", err.Error())
	}

	return nil
}

func parseTaskStatus(status string) (schemas.TaskStatus, *models.ErrorMessage) {
	switch schemas.TaskStatus(status) {
	case schemas.TaskPending, schemas.TaskInProgress, schemas.TaskCompleted:
		return schemas.TaskStatus(status), nil
	default:
		return "", models.NewErrorMessage("Status", "invalid")
	}
}

func taskFindError(err error) *models.ErrorMessage {
	if errors.Is(err, base.ErrNotFound) {
		return models.NewErrorMessage("Task", "Not found")
	}

	return models.NewErrorMessage("Database", err.Error())
}
