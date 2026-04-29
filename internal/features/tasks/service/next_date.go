package service

import (
	core_errors "n1ktarchik/go-final/internal/core/errors"
	"strconv"
	"strings"
	"time"
)

func (s *TasksService) NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", core_errors.RepetitionRule_Error()
	}

	startDate, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", core_errors.StartDate_Error()
	}

	parts := strings.Split(repeat, " ")
	rule := parts[0]

	switch rule {
	case "d":
		if len(parts) < 2 {
			return "", core_errors.Interval_D_Error("no interval is specified for the 'd' rule")
		}
		days, err := strconv.Atoi(parts[1])
		if err != nil || days < 1 || days > 400 {
			return "", core_errors.Interval_D_Error("incorrect interval for 'd' (1-400)")
		}

		current := startDate
		for {
			current = current.AddDate(0, 0, days)
			if current.After(now) {
				return current.Format("20060102"), nil
			}
		}

	case "y":
		current := startDate
		for {
			current = current.AddDate(1, 0, 0)
			if current.After(now) {
				return current.Format("20060102"), nil
			}
		}

	case "w":
		if len(parts) < 2 {
			return "", core_errors.Interval_W_Error("no days of the week are specified for 'w'")
		}
		weekdayStrs := strings.Split(parts[1], ",")
		weekdays := make(map[time.Weekday]bool)
		for _, s := range weekdayStrs {
			wd, err := strconv.Atoi(s)
			if err != nil || wd < 1 || wd > 7 {
				return "", core_errors.Interval_W_Error("incorrect day of the week")
			}
			if wd == 7 {
				weekdays[time.Sunday] = true
			} else {
				weekdays[time.Weekday(wd)] = true
			}
		}

		current := startDate
		for {
			current = current.AddDate(0, 0, 1)
			if current.After(now) && weekdays[current.Weekday()] {
				return current.Format("20060102"), nil
			}
		}

	case "m":
		if len(parts) < 2 {
			return "", core_errors.Interval_M_Error("no days of the month are specified for 'm'")
		}

		dayStrs := strings.Split(parts[1], ",")
		daysMap := make(map[int]bool)
		for _, s := range dayStrs {
			d, err := strconv.Atoi(s)
			if err != nil || d < -2 || d == 0 || d > 31 {
				return "", core_errors.Interval_M_Error("incorrect day of the month")
			}
			daysMap[d] = true
		}

		monthsMap := make(map[time.Month]bool)
		if len(parts) > 2 {
			monthStrs := strings.Split(parts[2], ",")
			for _, s := range monthStrs {
				m, err := strconv.Atoi(s)
				if err != nil || m < 1 || m > 12 {
					return "", core_errors.Interval_M_Error("incorrect month")
				}
				monthsMap[time.Month(m)] = true
			}
		}

		current := startDate
		for {
			current = current.AddDate(0, 0, 1)

			if len(monthsMap) > 0 && !monthsMap[current.Month()] {
				continue
			}

			day := current.Day()

			if daysMap[day] {
				if current.After(now) {
					return current.Format("20060102"), nil
				}
			}

			if daysMap[-1] {
				lastDay := lastDayOfMonth(current)
				if day == lastDay && current.After(now) {
					return current.Format("20060102"), nil
				}
			}

			if daysMap[-2] {
				lastDay := lastDayOfMonth(current)
				if day == lastDay-1 && current.After(now) {
					return current.Format("20060102"), nil
				}
			}
		}

	default:
		return "", core_errors.Rule_Error(rule)
	}
}

func lastDayOfMonth(t time.Time) int {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, t.Location()).Day()
}
