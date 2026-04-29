package http

import (
	resp "n1ktarchik/go-final/internal/core/transport/response"
	"net/http"
)

func (h *TasksHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	search := r.URL.Query().Get("search")

	tasks, err := h.tasksService.GetAllTasks(ctx, search)
	if err != nil {
		resp.RespondWithError(w, err)
		return
	}

	resp.RespondWithVallue(w, http.StatusOK, "tasks", tasks)
}

func (h *TasksHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")

	task, err := h.tasksService.GetTask(ctx, id)
	if err != nil {
		resp.RespondWithError(w, err)
		return
	}

	resp.RespondWithJSON(w, http.StatusOK, task)
}
