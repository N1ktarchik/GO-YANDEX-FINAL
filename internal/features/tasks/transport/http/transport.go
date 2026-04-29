package http

import (
	"context"
	"n1ktarchik/go-final/internal/core/domain"
	"time"
)

type TasksHandler struct {
	tasksService TasksService
}

type TasksService interface {
	NextDate(now time.Time, dstart string, repeat string) (string, error)

	CreateTask(ctx context.Context, task *domain.Task) (string, error)
	GetAllTasks(ctx context.Context, search string) ([]domain.Task, error)
	GetTask(ctx context.Context, id string) (*domain.Task, error)
	UpdateTask(ctx context.Context, task *domain.Task) error
	DeleteTask(ctx context.Context, id string) error
	CompleteTask(ctx context.Context, id string) error
}

func NewTasksTransport(tasksService TasksService) *TasksHandler {
	return &TasksHandler{
		tasksService: tasksService,
	}
}
