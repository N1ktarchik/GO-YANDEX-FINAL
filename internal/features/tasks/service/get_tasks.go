package service

import (
	"context"
	"n1ktarchik/go-final/internal/core/domain"
	core_errors "n1ktarchik/go-final/internal/core/errors"

	"strconv"
	"time"
)

func (s *TasksService) GetAllTasks(ctx context.Context, search string) ([]domain.Task, error) {
	dateFlag := false
	cleanSearch := ""

	if search != "" {
		t, err := time.Parse("02.01.2006", search)

		if err == nil {
			dateFlag = true
			cleanSearch = t.Format("20060102")

		} else {
			cleanSearch = search
		}
	}

	tasks, err := s.tasksRepository.GetAllTasks(ctx, cleanSearch, dateFlag)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TasksService) GetTask(ctx context.Context, id string) (*domain.Task, error) {
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return nil, core_errors.BadRequest("ivalid task id")
	}

	if parsedId < 0 {
		return nil, core_errors.BadRequest("ivalid task id")
	}

	task, err := s.tasksRepository.GetTask(ctx, parsedId)
	if err != nil {
		return nil, err
	}

	return task, nil
}
