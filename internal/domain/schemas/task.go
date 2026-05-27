package schemas

import (
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/domain/models/dates"
	"gorm.io/gorm"
)

type TaskStatus string

const (
	TaskPending    TaskStatus = "Pendente"
	TaskCompleted  TaskStatus = "Concluido"
	TaskInProgress TaskStatus = "Em progresso"
)

type Task struct {
	gorm.Model
	Name        models.GoalName `gorm:"embedded"`
	Description string          `gorm:"default:'Descrição não informada'"`
	Status      TaskStatus      `gorm:"type:task_status;not null;default:'Pendente'"`
	Dates       dates.TaskDates `gorm:"embedded"`
	GoalID      uint            `gorm:"not null"`
	Goal        Goal            `gorm:"foreignKey:GoalID"`
}
