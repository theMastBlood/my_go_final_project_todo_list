package main

import (
	"log"

	"github.com/theMastBlood/my_go_final_project_todo_list/pkg/db"
	"github.com/theMastBlood/my_go_final_project_todo_list/pkg/server"
)

func main() {

	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}

	server.StartServer()

}
