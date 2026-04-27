package repository

import (
	"context"
	core_errors "n1ktarchik/go-final/internal/core/errors"
)

func (r *TasksRepository) DeleteTask(ctx context.Context, id int) error {
	query := `DELETE FROM scheduler WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return core_errors.ServerError("deleting task from database faild")
	}

	count, err := result.RowsAffected()
	if err != nil {
		return core_errors.ServerError("error to get rows affected")
	}

	if count == 0 {
		return core_errors.BadRequest("task not found")
	}

	return nil
}
