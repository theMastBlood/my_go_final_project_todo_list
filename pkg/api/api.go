package api

import (
	"net/http"
)

func Init() {
	http.HandleFunc("/api/nextdate", NextDateHandler)
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r)
	default:
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
