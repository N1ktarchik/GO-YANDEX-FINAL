package http

import (
	core_errors "n1ktarchik/go-final/internal/core/errors"
	resp "n1ktarchik/go-final/internal/core/transport/response"
	"net/http"
	"time"
)

func (h *TasksHandler) NextDate(w http.ResponseWriter, r *http.Request) {
	nowString := r.URL.Query().Get("now")
	date := r.URL.Query().Get("date")
	repeat := r.URL.Query().Get("repeat")

	var now time.Time
	var err error
	if nowString == "" {
		now = time.Now().UTC()
	} else {
		now, err = time.Parse("20060102", nowString)
		if err != nil {
			resp.RespondWithError(w, core_errors.BadRequest("error to parse users time"))
			return
		}
	}

	nextDate, err := h.tasksService.NextDate(now, date, repeat)
	if err != nil {
		resp.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))
}
