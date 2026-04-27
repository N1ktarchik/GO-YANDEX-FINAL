package repository

import (
	"context"
	"n1ktarchik/go-final/internal/core/domain"
	core_errors "n1ktarchik/go-final/internal/core/errors"
)

func (r *TasksRepository) UpdateTask(ctx context.Context, task *domain.Task, id int) error {
	query := `UPDATE scheduler 
        	SET date = $1, 
            title = $2, 
            comment = $3, 
            repeat = $4 
            WHERE id = $5`

	updated, err := r.db.ExecContext(ctx, query, task.Date, task.Title, task.Comment, task.Repeat, id)

	if err != nil {
		return core_errors.ServerError("error to update task in databse")
	}

	count, err := updated.RowsAffected()
	if err != nil {
		return core_errors.ServerError("error to get rows affected")
	}

	if count == 0 {
		return core_errors.BadRequest("task not found")
	}

	return nil
}
