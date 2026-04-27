package repository

import (
	"context"
	"database/sql"
	"errors"
	"n1ktarchik/go-final/internal/core/domain"
	core_errors "n1ktarchik/go-final/internal/core/errors"
	"time"
)

func (r *TasksRepository) GetAllTasks(ctx context.Context, search string, dateFlag bool) ([]domain.Task, error) {
	var query string
	noSearchParam := true

	if search != "" {

		noSearchParam = false

		if dateFlag {
			query = `SELECT id,date,title,comment,repeat
					FROM scheduler 
					WHERE date = $1 
					LIMIT 30`
		} else {
			query = `SELECT id,date,title,comment,repeat
						FROM scheduler 
						WHERE title LIKE $1 
						OR comment LIKE $1 
						ORDER BY date 
						LIMIT 30`

			search = "%" + search + "%"

		}
	} else {
		query = `SELECT id,date,title,comment,repeat 
				FROM scheduler 
				WHERE date >= $1
				ORDER BY date 
				LIMIT 30`
	}

	var rows *sql.Rows
	var err error

	if noSearchParam {
		now := time.Now().Format("20060102")
		rows, err = r.db.QueryContext(ctx, query, now)
		if err != nil {
			return nil, core_errors.ServerError("failed to query tasks")
		}
	} else {
		if dateFlag {
			rows, err = r.db.QueryContext(ctx, query, search)
		} else {
			rows, err = r.db.QueryContext(ctx, query, search, search)
		}

		if err != nil {
			return nil, core_errors.ServerError("failed to query tasks")
		}
	}

	defer rows.Close()
	tasks := make([]domain.Task, 0)

	for rows.Next() {
		var taskModel taskModel

		err := rows.Scan(
			&taskModel.id,
			&taskModel.Date,
			&taskModel.Title,
			&taskModel.Comment,
			&taskModel.Repeat)

		if err != nil {
			return nil, core_errors.ServerError("failed to scan sheduler row")
		}

		tasks = append(tasks, modelToDomain(taskModel))

	}

	if err = rows.Err(); err != nil {
		return nil, core_errors.ServerError("rows error")
	}

	return tasks, nil
}

func (r *TasksRepository) GetTask(ctx context.Context, id int) (*domain.Task, error) {
	query := `SELECT id,date,title,comment,repeat 
			FROM scheduler 
			WHERE id=$1`

	taskModel := taskModel{}

	row := r.db.QueryRowContext(ctx, query, id)

	if err := taskModel.scan(row); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core_errors.TaskNotFaund(id)
		}

		return nil, core_errors.ServerError("error to get task by id from databse")
	}

	domainTask := modelToDomain(taskModel)
	return &domainTask, nil
}
