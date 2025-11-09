package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var tasks []Task

const fileName = "tasks.json"

func loadTasks() {
	file, err := os.ReadFile(fileName)
	if err == nil {
		json.Unmarshal(file, &tasks)
	}
}

func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(fileName, data, 0644)
}

func addTask(name string) {
	id := len(tasks) + 1
	tasks = append(tasks, Task{ID: id, Name: name, Done: false})
	saveTasks()
	fmt.Println("Task added:", name)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		status := "Pending"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("[%s] %d. %s\n", status, task.ID, task.Name)
	}
}

func markDone(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			saveTasks()
			fmt.Println("Task marked done:", task.Name)
			return
		}
	}
	fmt.Println("Task not found.")
}

func main() {
	loadTasks()

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  add <task name>     - Add new task")
		fmt.Println("  list                - List all tasks")
		fmt.Println("  done <task id>      - Mark task as done")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task name.")
			return
		}
		addTask(os.Args[2])
	case "list":
		listTasks()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide task ID.")
			return
		}
		var id int
		fmt.Sscanf(os.Args[2], "%d", &id)
		markDone(id)
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
