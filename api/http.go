package api

import (
	"cli-todo/commands"
	"encoding/json"
	"fmt"
	"net/http"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
	var data commands.AddTaskData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		msg := "fail to write HTTP response: " + err.Error()
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response: " + err.Error())
			return
		}
		fmt.Println(msg)
		return
	}
	fmt.Println(data)
	commands.LoadTasks()
	err = commands.AddTask(&data)
	if err != nil {
		msg := "fail to write HTTP response: " + err.Error()
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response: " + err.Error())
			return
		}
		fmt.Println(msg)
		return
	}
	commands.SaveTask()
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	commands.LoadTasks()
	dataTasks, err := json.Marshal(commands.Tasks)
	if err != nil {
		msg := "fail to write HTTP response: " + err.Error()
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response: " + err.Error())
			return
		}
		fmt.Println(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(dataTasks)
	if err != nil {
		fmt.Println("fail to write HTTP response: " + err.Error())
		return
	}
}

func HTTPServer() {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/list", listHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
		return
	}
}
