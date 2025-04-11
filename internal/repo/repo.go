package repo

import (
	"context"
	"github.com/mjolner36/skillsrock_todo/internal/entity"
)

type TaskRepository interface {
	GetAllTasks(ctx context.Context) ([]entity.Task, error)
	GetTask(ctx context.Context, id int) (entity.Task, error)
	CreateTask(ctx context.Context, task entity.Task) error
	DeleteTask(ctx context.Context, id int) error
	UpdateTask(ctx context.Context, task entity.Task, id int) error
}
