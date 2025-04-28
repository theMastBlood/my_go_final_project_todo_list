package api

import (
	"net/http"

	"github.com/theMastBlood/my_go_final_project_todo_list/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	if id == "" {
		writeJson(w, map[string]string{"error": "task ID not specified"}, http.StatusBadRequest)
		return
	}

	if err := db.DeleteTask(id); err != nil {
		writeJson(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	writeJson(w, map[string]string{}, http.StatusOK)
}
