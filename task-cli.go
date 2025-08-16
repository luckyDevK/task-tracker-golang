package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Status string

const (
	Todo Status = "todo"
	InProgress Status = "in-progress"
	Done Status = "Done"
)


type Task struct {
    Id          int64  `json:"id"`
    Description string `json:"description"`
    Status      Status `json:"status"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
}



func NewTask(desc string, status Status) Task {
	nowFormatted := time.Now().Format(time.RFC3339)

	newTask := Task{
		Id:          time.Now().Unix(),
		Description: desc,
		Status:      status,
		CreatedAt:   nowFormatted,
		UpdatedAt:   nowFormatted,
	}

	fmt.Printf("Task added successfully ID: %v\n", newTask.Id)

	return newTask
}


// func (t *Task) UnmarshalJSON(data []byte) error  {
// 	type Alias Task

// 	aux := &struct {
//         ID        int64  `json:"id"`
//         CreatedAt string `json:"created_at"`
//         *Alias
//     }{
//         Alias: (*Alias)(t),
//     }

// 	if err := json.Unmarshal(data, aux); err != nil{
// 		return err
// 	}

// 	t.Id = aux.ID
// 	t.CreatedAt = aux.CreatedAt
//     return nil
// }

func storeToJsonFile(newTask Task)  {
	var allTask []Task
	
	file, err := os.ReadFile("tasks.json")


	if err == nil {
		json.Unmarshal(file, &allTask)
	}

	allTask = append(allTask, newTask)


	data, err := json.MarshalIndent(allTask, "", "  ")
	
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}
	err = os.WriteFile("tasks.json", data, 0644)

	if err != nil {
        log.Fatal(err)
    }
}



func main()  {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter task: ")
	desc, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error enter task: %v\n", err)
		return
	}

	desc = strings.TrimSpace(desc)
	newTask := NewTask(desc, InProgress)
	
	fmt.Println(newTask, "newTask")

	storeToJsonFile(newTask)
	

	// if err != nil {
	// 	fmt.Printf("Error creating file: %v\n", err)
	// }

	// defer file.Close()

	// _, err = file.Write(jsonData)
	// if err != nil {
	// 	fmt.Printf("Error writing to file: %v\n", err)
	// 	return
	// }

	// 	fmt.Println("Data successfully written to user.json")

}

