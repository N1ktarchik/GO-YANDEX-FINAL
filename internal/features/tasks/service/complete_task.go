package service

import (
	"context"
	"time"
)

func (s *TasksService) CompleteTask(ctx context.Context, id string) error {

	task, err := s.GetTask(ctx, id)
	if err != nil {
		return err
	}

	if task.Repeat == "" {

		if err := s.DeleteTask(ctx, id); err != nil {
			return err
		}

		return nil
	}

	newDate, err := s.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		return err
	}

	task.Date = newDate

	if err := s.UpdateTask(ctx, task); err != nil {
		return err
	}

	return nil
}
