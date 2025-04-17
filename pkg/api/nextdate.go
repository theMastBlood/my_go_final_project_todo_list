package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {

	if repeat == "" {
		return "", fmt.Errorf("Повтор не задан")
	}

	date, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", fmt.Errorf("Неверный формат даты: %s", dstart)
	}

	switch repeat[0] {

	case 'd':
		parts := strings.Split(repeat, " ")
		if len(parts) != 2 {
			return "", fmt.Errorf("Неверный формат даты: %s", repeat)
		}
		interval, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf("Неверный формат даты: %s", repeat)
		} else if interval < 1 || interval > 400 {
			return "", fmt.Errorf("Задан недопустимый интервал повторения: %s", interval)
		}
		nextDate := date.AddDate(0, 0, interval)
		for !nextDate.After(now) {
			nextDate = nextDate.AddDate(0, 0, interval)
		}
		return nextDate.Format(dateFormat), nil

	case 'y':
		nextDate := date.AddDate(1, 0, 0)
		for !nextDate.After(now) {
			nextDate = nextDate.AddDate(1, 0, 0)
		}
		return nextDate.Format(dateFormat), nil

	default:
		return "", fmt.Errorf("Неподдерживаемый формат даты: %s", repeat)
	}
}

func NextDateHandler(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	nowFormat, err := time.Parse(dateFormat, now)
	if err != nil {
		http.Error(w, "Неверный формат даты", http.StatusBadRequest)
		return
	}

	nextDate, err := NextDate(nowFormat, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(nextDate))
}
