package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)

// Task represents a task item
type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// ---------------------- Utility Functions ----------------------

// getCurrentTime returns formatted current time
func getCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}

// trimSpace trims spaces from input string
func trimSpace(text string) string {
	return strings.TrimSpace(text)
}

// stringIdToInt converts a string ID to int, exits on error
func stringIdToInt(idStr string) int {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. ID must be a number.")
		os.Exit(1)
	}
	return id
}

// ---------------------- File Operations ----------------------

// loadTasks reads tasks from "tasks.json"
func loadTasks() []Task {
	var allTasks []Task
	file, err := os.ReadFile("tasks.json")
	if err == nil {
		json.Unmarshal(file, &allTasks)
	}
	return allTasks
}

// saveTasks writes tasks to "tasks.json"
func saveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}
	if err := os.WriteFile("tasks.json", data, 0644); err != nil {
		log.Fatal(err)
	}
}

// ---------------------- Task Operations ----------------------

// newTask creates a new task and appends it to the list
func newTask(desc string) []Task {
	tasks := loadTasks()
	task := Task{
		Id:          len(tasks) + 1,
		Description: desc,
		Status:      Todo,
		CreatedAt:   getCurrentTime(),
		UpdatedAt:   getCurrentTime(),
	}
	fmt.Printf("Task added successfully ID: %v\n", task.Id)
	tasks = append(tasks, task)
	return tasks
}

// getTaskFromId returns a pointer to the task with matching id
func getTaskFromId(id int, tasks []Task) (*Task, bool) {
	for i, t := range tasks {
		if t.Id == id {
			return &tasks[i], true
		}
	}
	return nil, false
}

// updateTask updates the description of a task
func updateTask(id int, desc string) {
	tasks := loadTasks()
	task, found := getTaskFromId(id, tasks)
	if !found {
		fmt.Println("Task not found")
		return
	}
	task.Description = desc
	task.UpdatedAt = getCurrentTime()
	saveTasks(tasks)
	fmt.Println("Task updated successfully")
}

// markTaskProgress updates task status (done or in-progress)
func markTaskProgress(id int, status Status) {
	tasks := loadTasks()
	task, found := getTaskFromId(id, tasks)
	if !found {
		fmt.Println("Task not found")
		return
	}
	task.Status = status
	task.UpdatedAt = getCurrentTime()
	saveTasks(tasks)
	fmt.Printf("Task %v marked as %s\n", id, status)
}

// ---------------------- Display Functions ----------------------

// displayTasks prints a list of tasks
func displayTasks(tasks []Task) {
	for _, t := range tasks {
		fmt.Printf("ID: %d | %s | %s | Created: %s | Updated: %s\n",
			t.Id, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
	}
}

// filterTasksByStatus returns tasks matching a status
func filterTasksByStatus(status Status) []Task {
	tasks := loadTasks()
	var filtered []Task
	for _, t := range tasks {
		if t.Status == status {
			filtered = append(filtered, t)
		}
	}
	return filtered
}

// ---------------------- Command Handlers ----------------------

func handleListCommand(args []string) {
	if len(args) == 1 {
		displayTasks(loadTasks())
		return
	}
	filter := args[1]
	switch filter {
	case "done":
		displayTasks(filterTasksByStatus(Done))
	case "todo":
		displayTasks(filterTasksByStatus(Todo))
	case "in-progress":
		displayTasks(filterTasksByStatus(InProgress))
	default:
		fmt.Println("Unknown list filter:", filter)
	}
}

func handleCommandInput() {
	args := os.Args[1:] // skip program name
	if len(args) < 1 {
		fmt.Println("Please provide a command.")
		return
	}

	command := args[0]
	switch command {
	case "new":
		if len(args) < 2 {
			fmt.Println("Usage: task-cli new \"task description\"")
			return
		}
		tasks := newTask(trimSpace(args[1]))
		saveTasks(tasks)

	case "update":
		if len(args) < 3 {
			fmt.Println("Usage: task-cli update (id) \"updated description\"")
			return
		}
		id := stringIdToInt(args[1])
		desc := trimSpace(args[2])
		updateTask(id, desc)

	case "mark-in-progress":
		if len(args) < 2 {
			fmt.Println("Usage: task-cli mark-in-progress (id)")
			return
		}
		id := stringIdToInt(args[1])
		markTaskProgress(id, InProgress)

	case "mark-done":
		if len(args) < 2 {
			fmt.Println("Usage: task-cli mark-done (id)")
			return
		}
		id := stringIdToInt(args[1])
		markTaskProgress(id, Done)

	case "list":
		handleListCommand(args)

	default:
		fmt.Println("Unknown command:", command)
	}
}

// ---------------------- Main ----------------------

func main() {
	handleCommandInput()
}
