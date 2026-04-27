package repository

import (
	"context"
	"n1ktarchik/go-final/internal/core/domain"
	core_errors "n1ktarchik/go-final/internal/core/errors"
)

func (r *TasksRepository) CreateTask(ctx context.Context, task *domain.Task) (int, error) {
	query := `INSERT INTO scheduler (date,title,comment,repeat) 
			VALUES ($1,$2,$3,$4)
			RETURNING id`

	var id int
	if err := r.db.QueryRowContext(ctx, query, task.Date, task.Title, task.Comment, task.Repeat).Scan(&id); err != nil {
		return 0, core_errors.ServerError("error to save task in database")
	}

	return id, nil
}
