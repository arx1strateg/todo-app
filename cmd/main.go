package cmd

import (
	"github.com/arx1strateg/todo-app"
	"github.com/arx1strateg/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	server := new(todo.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
