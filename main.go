package main

import (
	"fmt"
	"os"
	"cli-todo/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: todo add <текст задачи>")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		commands.LoadTasks()
		commands.AddTask(os.Args[2:])
		commands.SaveTask()
	case "list":
		commands.LoadTasks()
    	commands.ListTasks()
	case "done":
		commands.DoneTask(os.Args[2:])
	case "delete":
		commands.DeleteTask(os.Args[2:])
	}
}