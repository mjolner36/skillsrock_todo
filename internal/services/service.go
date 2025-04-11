package services

import (
	"context"
	"github.com/mjolner36/skillsrock_todo/internal/entity"
	"github.com/mjolner36/skillsrock_todo/internal/repo"
	"time"
)

type Service struct {
	repo repo.TaskRepository
}

func NewService(repo repo.TaskRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllTasks(ctx context.Context) ([]entity.Task, error) {
	return s.repo.GetAllTasks(ctx)
}

func (s *Service) GetTask(ctx context.Context, id int) (entity.Task, error) {
	return s.repo.GetTask(ctx, id)
}

func (s *Service) CreateTask(ctx context.Context, task entity.Task) error {
	task.CreatedAt = time.Now()
	return s.repo.CreateTask(ctx, task)
}

func (s *Service) DeleteTask(ctx context.Context, id int) error {
	return s.repo.DeleteTask(ctx, id)
}

func (s *Service) UpdateTask(ctx context.Context, newTask entity.Task, id int) error {
	oldTask, err := s.repo.GetTask(ctx, id)
	if err != nil {
		return err
	}

	newTask.UpdatedAt = time.Now()
	newTask.Status = oldTask.Status
	newTask.Description = oldTask.Description
	newTask.Title = oldTask.Title

	return s.repo.UpdateTask(ctx, newTask, id)
}
