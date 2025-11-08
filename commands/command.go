package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type AddTaskData struct {
	Text string `json:"text"`
}
type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var Tasks []Task

func AddTask(a *AddTaskData) error {
	if a.Text == "" {
		return errors.New("empty text")
	}

	task := Task{
		ID:   len(Tasks) + 1,
		Text: a.Text,
		Done: false,
	}
	Tasks = append(Tasks, task)
	fmt.Printf("Добавлено: %d. [ ] %s\n", len(Tasks), a.Text)
	return nil
}

func ListTasks() {
	if len(Tasks) == 0 {
		fmt.Println("Список задач пуст")
		return
	}

	for i, task := range Tasks {
		if task.Done {
			fmt.Printf("%d. [x] %s\n", i+1, task.Text)
		} else {
			fmt.Printf("%d. [ ] %s\n", i+1, task.Text)
		}
	}
}

func LoadTasks() {
	_, err := os.Stat("tasks.json")
	if err != nil {
		fmt.Println("Файл не найден")
		return
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

func DoneTask(args []string) {
	LoadTasks()
	if len(args) == 0 {
		fmt.Println("Нужно ввести номер задачи")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Номер задачи должен быть числом")
		return
	}

	for i, task := range Tasks {
		if task.ID == id {
			Tasks[i].Done = true
			fmt.Printf("Задача %d отмечена как выполненная: %s\n", id, task.Text)
			break
		}
	}
	SaveTask()
}

func DeleteTask(args []string) {
	LoadTasks()
	if len(args) == 0 {
		fmt.Println("Нужно ввести номер задачи")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Номер задачи должен быть числом")
		return
	}

	for i, task := range Tasks {
		if task.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			fmt.Printf("Задача %d удалена: %s\n", id, task.Text)
			break
		}
	}
	SaveTask()
}
