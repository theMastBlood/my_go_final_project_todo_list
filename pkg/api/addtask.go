package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/theMastBlood/my_go_final_project_todo_list/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeJson(w, map[string]string{"error": "error reading request body"}, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &task); err != nil {
		writeJson(w, map[string]string{"error": "JSON deserialization error"}, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeJson(w, map[string]string{"error": "task title not specified"}, http.StatusBadRequest)
		return
	}

	if err := checkDate(&task); err != nil {
		writeJson(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeJson(w, map[string]string{"error": "error adding task to database"}, http.StatusInternalServerError)
		return
	}

	writeJson(w, map[string]string{"id": strconv.FormatInt(id, 10)}, http.StatusOK)
}

func checkDate(task *db.Task) error {
	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format(dateFormat)
		return nil
	}

	t, err := time.Parse(dateFormat, task.Date)
	if err != nil {
		return fmt.Errorf("date is not in the correct format, expected: %s", dateFormat)
	}

	if t.Before(now.Truncate(24 * time.Hour)) {
		task.Date = now.Format(dateFormat)
	}

	if err := validateTask(task); err != nil {
		return err
	}

	return nil
}

func writeJson(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "error while serializing data", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "error sending response", http.StatusInternalServerError)
	}
}
