package main

import (
	"fmt"
	"os"
	"strings"
	"encoding/json"
)

type Task struct {
	ID   int
	Text string
	Done bool
}

var Tasks []Task

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: todo add <текст задачи>")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		LoadTasks()
		addTask(os.Args[2:])
		SaveTask()
	case "list":
		LoadTasks()
    	listTasks()
	}
}

func addTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Нужно ввести текст задачи")
		return
	}

	text := strings.Join(args, " ")
	task := Task{
		ID:   len(Tasks) + 1,
		Text: text,
		Done: false,
	}
	Tasks = append(Tasks, task)
	fmt.Printf("Добавлено: %d. [ ] %s\n", len(Tasks), text)
}

func listTasks() {
	if len(Tasks) == 0 {
		fmt.Println("Список задач пуст")
	}

	for i, task := range Tasks {
		fmt.Printf("%d. [ ] %s\n", i+1, task.Text)
	}
}

func LoadTasks() {
	_, err := os.Stat("tasks.json")
	if err != nil {
		fmt.Println("Файл не найден")
	} else {
		data, _ := os.ReadFile("tasks.json")
		_ = json.Unmarshal(data, &Tasks)
	}
}

func SaveTask() {
	jsonData, err := json.MarshalIndent(Tasks, "", " ")
	if err != nil {
		fmt.Println("Error:", err)
        return
	}
	_ = os.WriteFile("tasks.json", jsonData, 0644)
}