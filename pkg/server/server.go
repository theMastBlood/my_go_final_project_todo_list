package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort = "7540"
	webDir      = "./web"
)

func StartServer() {
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
