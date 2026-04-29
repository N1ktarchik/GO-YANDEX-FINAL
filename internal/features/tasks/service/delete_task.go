package service

import (
	"context"
	core_errors "n1ktarchik/go-final/internal/core/errors"
	"strconv"
)

func (s *TasksService) DeleteTask(ctx context.Context, id string) error {
	validID, err := strconv.Atoi(id)

	if err != nil {
		return core_errors.BadRequest("parse task id faild")
	}

	if validID < 0 {
		return core_errors.BadRequest("invalid task id")
	}

	if err := s.tasksRepository.DeleteTask(ctx, validID); err != nil {
		return err
	}

	return nil
}
