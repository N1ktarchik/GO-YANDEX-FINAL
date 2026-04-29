package http

import (
	req "n1ktarchik/go-final/internal/core/transport/request"
	resp "n1ktarchik/go-final/internal/core/transport/response"
	"net/http"
)

func (h *TasksHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	task := &TaskDto{}

	if err := req.DecodeJSON(r, task); err != nil {
		resp.RespondWithError(w, err)
		return
	}

	if err := h.tasksService.UpdateTask(ctx, task.Todomain()); err != nil {
		resp.RespondWithError(w, err)
		return
	}

	resp.RespondWithJSON(w, http.StatusOK, map[string]any{})
}
