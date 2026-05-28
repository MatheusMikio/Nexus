package service

import (
	"github.com/MatheusMikio/Nexus/internal/domain/dtos/task"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/repository"
	"github.com/google/uuid"
)

type ITaskService interface {
	GetAllTasks(parameters parameters.PaginationQuery, goalId uint, userId uuid.UUID) ([]*task.Response, *models.ErrorMessage)
	GetById(id uint, userId uuid.UUID) (*task.Response, *models.ErrorMessage)
	Create(goalID uint, task *task.Request, userId uuid.UUID) []*models.ErrorMessage
	Update(id uint, task *task.Update, userId uuid.UUID) []*models.ErrorMessage
	Delete(id uint, userId uuid.UUID) *models.ErrorMessage
}

type TaskService struct {
	TaskRepository repository.ITaskRepository
	UserRepository repository.IUserRepository
}

func NewTaskService(taskRepo repository.ITaskRepository, userRepo repository.IUserRepository) ITaskService {
	return &TaskService{
		TaskRepository: taskRepo,
		UserRepository: userRepo,
	}
}

func (t *TaskService) GetAllTasks(parameters parameters.PaginationQuery, goalId uint, userId uuid.UUID) ([]*task.Response, *models.ErrorMessage) {
	panic("unimplemented")
}

func (t *TaskService) GetById(id uint, userId uuid.UUID) (*task.Response, *models.ErrorMessage) {
	panic("unimplemented")
}

func (t *TaskService) Create(goalID uint, task *task.Request, userId uuid.UUID) []*models.ErrorMessage {
	panic("unimplemented")
}

func (t *TaskService) Update(id uint, task *task.Update, userId uuid.UUID) []*models.ErrorMessage {
	panic("unimplemented")
}

func (t *TaskService) Delete(id uint, userId uuid.UUID) *models.ErrorMessage {
	panic("unimplemented")
}
