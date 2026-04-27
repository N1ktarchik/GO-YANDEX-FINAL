package service

import (
	"context"
	"n1ktarchik/go-final/internal/core/domain"
)

type TasksService struct {
	tasksRepository TasksRepository
}

type TasksRepository interface {
	CreateTask(ctx context.Context, task *domain.Task) (int, error)
	GetAllTasks(ctx context.Context, search string, dateFlag bool) ([]domain.Task, error)
	GetTask(ctx context.Context, id int) (*domain.Task, error)
	UpdateTask(ctx context.Context, task *domain.Task, id int) error
	DeleteTask(ctx context.Context, id int) error
}

func NewTasksService(tasksRepository TasksRepository) *TasksService {
	return &TasksService{
		tasksRepository: tasksRepository,
	}
}
