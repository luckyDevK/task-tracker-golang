
# Task Tracker CLI

Task Tracker is a simple command-line interface (CLI) application built in Go to help you manage your tasks. You can add, update, delete, and track the progress of your tasks directly from the terminal.

---

## Features

- Add new tasks
- Update existing tasks
- Mark tasks as "in-progress" or "done"
- List all tasks
- List tasks filtered by status (`todo`, `in-progress`, `done`)
- Tasks are stored in a JSON file (`tasks.json`) in the current directory

---

## Getting Started

### Prerequisites

- Go installed (1.20+ recommended)
- Terminal or command prompt

### Running the CLI:

1.  Clone the repository or copy your code into a folder.
2. Navigate to the project folder in the terminal.
3. Run the CLI with:

```bash
go run main.go <command> [arguments]
```
### **Commands**

#### Add a New Task

```bash
go run main.go new "Task description"
```


#### Example:

```bash
go run main.go new "Buy groceries"

```

###

#### Update an Existing Task

```bash
go run main.go update <id> "Updated task description"

```

#### Example:

```bash
go run main.go update 1 "Buy groceries and cook dinner"
```
###

#### Delete a Task:

```bash
go run main.go delete <id>
```

#### Example:

```bash
go run main.go delete 1
```

###

#### Mark a Task as In Progress

```bash
go run main.go mark-in-progress <id>
```

#### Example:

```bash
go run main.go mark-in-progress 2
```

##

#### Mark a Task as Done

```bash
go run main.go mark-done <id>
```

#### Example
```bash
go run main.go mark-done 2
```
##

#### List All Tasks

```bash
go run main.go list
```
##

#### List Tasks by Status

```bash
go run main.go list todo
go run main.go list in-progress
go run main.go list done
```

##

## Task Structure

Each task has the following properties:

- id: Unique identifier for the task
- description: Short description of the task
- status: Task status (todo, in-progress, done)
- created_at: Timestamp when the task was created
- updated_at: Timestamp when the task was last updated

##

## Notes

- All tasks are stored in a tasks.json file in the current working directory.
- If ``tasks.json`` does not exist, it will be created automatically.
- Commands are case-sensitive, and task IDs must be numbers.

## 

### Example Workflow:

```bash
# Add a new task
go run main.go new "Finish homework"

# Update a task
go run main.go update 1 "Finish math homework"

# Mark task as in-progress
go run main.go mark-in-progress 1

# Mark task as done
go run main.go mark-done 1

# List all tasks
go run main.go list

# List tasks by status
go run main.go list todo
go run main.go list done

```
