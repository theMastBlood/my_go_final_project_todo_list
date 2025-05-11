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
		writeJson(w, map[string]string{"error": "JSON deserialization error"}, http.StatusBadRequest)
		return
	}

	if task.ID == "" {
		writeJson(w, map[string]string{"error": "task ID not specified"}, http.StatusBadRequest)
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
		return fmt.Errorf("no date specified")
	}

	t, err := time.Parse(dateFormat, task.Date)
	if err != nil {
		return fmt.Errorf("date is not in the correct format, expected: %s", dateFormat)
	}

	if t.Truncate(24 * time.Hour).Before(now.Truncate(24 * time.Hour)) {
		return fmt.Errorf("date cannot be earlier than current")
	}

	if task.Repeat != "" {
		parts := strings.Split(task.Repeat, " ")
		if parts[0] == "d" {
			if len(parts) != 2 {
				return fmt.Errorf("invalid repeat format: %s", task.Repeat)
			}
			interval, err := strconv.Atoi(parts[1])
			if err != nil {
				return fmt.Errorf("invalid repeat interval specified: %s", task.Repeat)
			} else if interval < 1 || interval > 400 {
				return fmt.Errorf("invalid repeat interval specified: %d", interval)
			}
		} else if parts[0] == "y" {
		} else {
			return fmt.Errorf("unsupported repeat format: %s", task.Repeat)
		}
	}
	return nil
}
