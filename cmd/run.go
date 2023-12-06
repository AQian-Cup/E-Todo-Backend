package main

import (
	eTodoBackend "e-todo-backend/pkg/e-todo-backend"
	"os"
)

func main() {
	command := eTodoBackend.GetCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
