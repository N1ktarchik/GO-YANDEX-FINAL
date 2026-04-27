package service

import (
	"context"
	"n1ktarchik/go-final/internal/core/domain"
	core_errors "n1ktarchik/go-final/internal/core/errors"
	"strconv"
	"time"
)

func (s *TasksService) UpdateTask(ctx context.Context, task *domain.Task) error {
	parsedId, err := strconv.Atoi(task.Id)
	if err != nil {
		return core_errors.BadRequest("ivalid task id")
	}

	if parsedId < 0 {
		return core_errors.BadRequest("ivalid task id")
	}

	if task.Title == "" {
		return core_errors.BadRequest("length of tasks title can not be zero ")
	}

	now := time.Now()
	today := now.Format("20060102")

	if task.Date == "" {
		task.Date = today
	}

	if _, err := time.Parse("20060102", task.Date); err != nil {
		return core_errors.BadRequest("error to parse tasks date")
	}

	if task.Repeat != "" {

		nextDate, err := s.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return core_errors.BadRequest("invalid repetition rule")
		}

		if task.Date < today {
			task.Date = nextDate
		}

	} else {

		if task.Date < today {
			task.Date = today
		}
	}

	if err := s.tasksRepository.UpdateTask(ctx, task, parsedId); err != nil {
		return err
	}

	return nil

}
