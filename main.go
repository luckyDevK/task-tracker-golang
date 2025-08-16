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
	Todo Status = "todo"
	InProgress Status = "in-progress"
	Done Status = "done"
)


type Task struct {
    Id          int  `json:"id"`
    Description string `json:"description"`
    Status      Status `json:"status"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
}

var currentTimeFormatted string = time.Now().Format(time.RFC3339) 


func newTask(desc string) []Task {
    nowFormatted := time.Now().Format(time.RFC3339)

	tasks := readTasks()

    task := Task{
        Id:          len(tasks) + 1, 
        Description: desc,
        Status:      Todo,
        CreatedAt:   nowFormatted,
        UpdatedAt:   currentTimeFormatted,
    }

    fmt.Printf("Task added successfully ID: %v\n", task.Id)

    tasks = append(tasks, task)

    return tasks
}

func getTaskFromId(id int, tasks []Task) (*Task, bool)  {
    for i, t := range tasks {
        if t.Id == id {
           return &tasks[i], true
    }
}

return nil, false
}

func (t *Task) MarkDone()  {
    t.Status = Done
    t.UpdatedAt = currentTimeFormatted
}

func (t *Task) MarkInProgress()  {
    t.Status = InProgress
    t.UpdatedAt = currentTimeFormatted
}



func updateTask(id int, desc string)  {
    tasks := readTasks()

    task, found := getTaskFromId(id, tasks)

    if found {
        task.Description = desc
        task.UpdatedAt = currentTimeFormatted
        saveTasks(tasks)
        fmt.Println("task updated")
    }else{
        fmt.Println("task not found")
    }
}

func markTaskProgres(id int, progress string)  {
    tasks := readTasks()

    task, found := getTaskFromId(id, tasks)

    if !found {
        fmt.Println("not item matched!!")
        return
    }

    switch progress {
    case "mark-in-progress":
    task.MarkInProgress()
    fmt.Printf("Successfully marked %v as in-progress", id)
    case "mark-done":
    task.MarkDone()
    fmt.Printf("Successfully marked %v as done\n", id)
    default:
    fmt.Println("Invalid progress command")
    return
}


    saveTasks(tasks)
}

func displayTasks(tasks []Task)  {
     for _, t := range tasks {
    fmt.Printf("ID: %d | %s | %s | Created: %s | Updated: %s\n",
        t.Id, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
}
}

func allTasks()  {
    tasks := readTasks()

   displayTasks(tasks)

}

func doneTasks()  {
    tasks := readTasks()

    var done []Task

    for _, t := range tasks {
        if t.Status == Done {
          done = append(done, t)
        }
    }
    fmt.Println("triggered")

    displayTasks(done)
}





// func updateTask(id int64)  {
	
// }


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

func readTasks() []Task  {
    var allTasks []Task
	
	file, err := os.ReadFile("tasks.json")

	if err == nil {
		json.Unmarshal(file, &allTasks)
	}

    
    return allTasks
}

func saveTasks(tasks []Task)  {
	data, err := json.MarshalIndent(tasks, "", "  ")
	
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}
	err = os.WriteFile("tasks.json", data, 0644)

	if err != nil {
        log.Fatal(err)
    }
}

func trimSpace(text string) string  {
    return strings.TrimSpace(text)
}

func stringIdToInt(idStr string) int {
   id, err := strconv.Atoi(idStr)

   if err != nil {
		fmt.Println("Invalid ID. ID must be a number.")
		os.Exit(1) 
	}

    return id
}

func handleCommanInput()  {
	  args := os.Args[1:] // skip program name

  
    if len(args) < 1 {
        fmt.Println("Please provide a command.")
        return
    }

    command := args[0] // first argument is the command

	
    
    
	switch command {
    case "new":
        if len(args) < 2 {
            fmt.Println("Usage: task-cli new \"task description\"")
            return
        }
         
         desc := trimSpace(args[1])
          tasks := newTask(desc)
          
            saveTasks(tasks)
        
    case "update":
        if len(args) < 3 {
            fmt.Println("Usage: task-cli update (id) \"updated description\"")
            return
        }

        id := stringIdToInt(args[1])
        desc := trimSpace(args[2])

        updateTask(id, desc)

        fmt.Printf("id %v desc %v\n", id, desc)

    case "mark-in-progress" : 
        if len(args) < 2 {
            fmt.Println("Usage: task-cli mark-in-progress (id)")
            return
        }
        
         id := stringIdToInt(args[1])
        progress := args[0]

        markTaskProgres(id, progress)

    case "mark-done" :
        if len(args) < 2 {
            fmt.Println("Usage: task-cli mark-done (id)")
            return
        }

          id := stringIdToInt(args[1])
        progress := args[0]

        markTaskProgres(id, progress)

    case "list" : 
        // if len(args) > 1 {
        //     fmt.Println("Usage: task-cli list")
        //     return
        // }

    // allTasks()

    case "list done" :
         if len(args) > 2 {
            fmt.Println("Usage: task-cli list done")
            return
         }
    
          fmt.Println("triggered")
    doneTasks()
    


	default:
        fmt.Println("Unknown command:", command)
    
	} 
}



func main()  {

	handleCommanInput()

	// reader := bufio.NewReader(os.Stdin)

	// fmt.Print("task-cli add ")
	// desc, err := reader.ReadString('\n')

	// if err != nil {
	// 	fmt.Printf("Error enter task: %v\n", err)
	// 	return
	// }

	// desc = strings.TrimSpace(desc)

	
}

