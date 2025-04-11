package repo

import (
	"context"
	"github.com/mjolner36/skillsrock_todo/internal/entity"
	"sync"
)

type InMemoryRepo struct {
	mu     sync.RWMutex
	tasks  map[int]entity.Task
	nextID int
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		tasks:  make(map[int]entity.Task),
		nextID: 0,
	}
}

func (r *InMemoryRepo) GetAllTasks(ctx context.Context) ([]entity.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]entity.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *InMemoryRepo) GetTask(ctx context.Context, id int) (entity.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task := r.tasks[id]
	return task, nil
}

func (r *InMemoryRepo) CreateTask(ctx context.Context, task entity.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	task.ID = r.nextID
	r.tasks[r.nextID] = task
	r.nextID++
	return nil
}

func (r *InMemoryRepo) DeleteTask(ctx context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.tasks, id)
	return nil
}

func (r *InMemoryRepo) UpdateTask(ctx context.Context, task entity.Task, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[id] = task
	return nil
}
