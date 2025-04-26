package api

import (
	"net/http"

	"github.com/theMastBlood/my_go_final_project_todo_list/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(10)
	if err != nil {
		writeJson(w, map[string]string{"error": "ошибка при получении списка задач"}, http.StatusInternalServerError)
		return
	}

	writeJson(w, TasksResp{
		Tasks: tasks,
	}, http.StatusOK)
}
