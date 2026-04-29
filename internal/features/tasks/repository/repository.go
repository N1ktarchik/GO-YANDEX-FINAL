package repository

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type TasksRepository struct {
	db *sql.DB
}

func NewTasksRepository(dbFile string) (*TasksRepository, error) {
	repo := &TasksRepository{}

	if err := repo.init(dbFile); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *TasksRepository) init(dbFile string) error {
	var err error
	r.db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	if err = r.db.Ping(); err != nil {
		return err
	}

	if _, err = r.db.Exec(schema); err != nil {
		return err
	}

	return nil
}
