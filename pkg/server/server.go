package server

import (
	"log"
	"net/http"
	"os"

	"github.com/theMastBlood/my_go_final_project_todo_list/pkg/api"
)

const defaultPort = "7540"

func StartServer() {

	api.Init()

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = defaultPort
	}

	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatalf("server starting error: %v\n", err)
		}
	}()

	log.Printf("Сервер запущен на порте: " + port)

	select {}
}
