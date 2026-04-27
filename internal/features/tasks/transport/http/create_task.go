package http

import (
	req "n1ktarchik/go-final/internal/core/transport/request"
	resp "n1ktarchik/go-final/internal/core/transport/response"
	"net/http"
)

func (h *TasksHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	task := &TaskDto{}

	if err := req.DecodeJSON(r, task); err != nil {
		resp.RespondWithError(w, err)
		return
	}

	id, err := h.tasksService.CreateTask(ctx, task.Todomain())
	if err != nil {
		resp.RespondWithError(w, err)
		return
	}

	resp.RespondWithVallue(w, http.StatusCreated, "id", id)
}
