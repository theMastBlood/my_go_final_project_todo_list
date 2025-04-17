package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/theMastBlood/my_go_final_project_todo_list/pkg/api"
)

const (
	defaultPort = "7540"
	webDir      = "./web"
)

func StartServer() {

	api.Init()

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatalf("Ошибка запуска сервера: %v\n", err)
		}
	}()

	fmt.Println("Сервер запущен на порте: " + port)

	select {}
}
