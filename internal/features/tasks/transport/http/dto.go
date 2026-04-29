package http

import "n1ktarchik/go-final/internal/core/domain"

type TaskDto struct {
	Id      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func (t *TaskDto) Todomain() *domain.Task {
	return &domain.Task{
		Id:      t.Id,
		Date:    t.Date,
		Title:   t.Title,
		Comment: t.Comment,
		Repeat:  t.Repeat,
	}
}
