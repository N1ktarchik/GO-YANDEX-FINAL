package http

import (
	resp "n1ktarchik/go-final/internal/core/transport/response"
	"net/http"
)

func (h *TasksHandler) CompleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")

	if err := h.tasksService.CompleteTask(ctx, id); err != nil {
		resp.RespondWithError(w, err)
		return
	}

	resp.RespondWithJSON(w, http.StatusOK, map[string]any{})
}
