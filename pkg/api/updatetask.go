package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/theMastBlood/my_go_final_project_todo_list/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJson(w, map[string]string{"error": "ошибка десериализации JSON"}, http.StatusBadRequest)
		return
	}

	if task.ID == "" {
		writeJson(w, map[string]string{"error": "не указан идентификатор задачи"}, http.StatusBadRequest)
		return
	}

	if err := validateTask(&task); err != nil {
		writeJson(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		writeJson(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	writeJson(w, map[string]string{}, http.StatusOK)
}

func validateTask(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		return fmt.Errorf("не указана дата")
	}

	t, err := time.Parse(dateFormat, task.Date)
	if err != nil {
		return fmt.Errorf("дата представлена в неправильном формате, ожидается: %s", dateFormat)
	}

	if t.Before(now) {
		return fmt.Errorf("дата не может быть раньше текущей")
	}

	if task.Repeat != "" {
		parts := strings.Split(task.Repeat, " ")
		if parts[0] == "d" {
			if len(parts) != 2 {
				return fmt.Errorf("неверный формат даты: %s", task.Repeat)
			}
			interval, err := strconv.Atoi(parts[1])
			if err != nil {
				return fmt.Errorf("неверный формат даты: %s", task.Repeat)
			} else if interval < 1 || interval > 400 {
				return fmt.Errorf("задан недопустимый интервал повторения: %d", interval)
			}
		}
	}

	return nil
}
