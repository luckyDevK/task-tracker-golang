package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Status string

const (
	Todo Status = "todo"
	InProgress Status = "in-progress"
	Done Status = "Done"
)


type Task struct {
	id int64
	Description string `description:"desc"`
	Status Status `json:"status"`
	createdAt time.Time 
	UpdatedAt time.Time `json:"updated_at"`
}




func NewTask(desc string, status Status) Task  {
	now := time.Now()

	newTask := Task{
		id: time.Now().UnixNano(),
		Description: desc,
		Status: status,
		createdAt: now,
		UpdatedAt: now,
	}

	return newTask
}

func (t Task) MarshalJSON() ([]byte, error)  {
	type Alias Task

	return json.Marshal(&struct {
		ID int64 `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		Alias
	}{
		ID: t.id,
		CreatedAt: t.createdAt,
		Alias: (Alias)(t),
	})
}



func main()  {
	taskOne := NewTask("solo", Todo)
	

	data, _ := json.MarshalIndent(taskOne, "", "  ")

	
	fmt.Println(string(data))
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("error when enter u name")
	}

	fmt.Println("Hello", name)

	// user := User{
	// 	Name:  "Jane Smith",
	// 	Email: "jane.smith@example.com",
	// 	Age:   25,
	// }

	// jsonData, err := json.MarshalIndent(user, "", "  ")
	// if err != nil {
	// 	fmt.Printf("Error marshaling JSON: %v\n", err)
	// 	return
	// }

	// file, err := os.Create("user.json")
	

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

