package repository

import (
	"database/sql"
	"n1ktarchik/go-final/internal/core/domain"
	"strconv"
)

type taskModel struct {
	id      int64
	Date    string
	Title   string
	Comment string
	Repeat  string
}

func (m *taskModel) scan(row *sql.Row) error {
	return row.Scan(
		&m.id,
		&m.Date,
		&m.Title,
		&m.Comment,
		&m.Repeat,
	)
}

func modelToDomain(model taskModel) domain.Task {
	return domain.Task{
		Id:      strconv.Itoa(int(model.id)),
		Date:    model.Date,
		Title:   model.Title,
		Comment: model.Comment,
		Repeat:  model.Repeat,
	}
}
