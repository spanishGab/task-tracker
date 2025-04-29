# task-tracker

Task tracker lets you manage your daily tasks. It was built with the intention to learn more about the hexagonal architecture and the base project requirements were extracted from [Roadmap.sh](https://roadmap.sh/projects/task-tracker).

## How to build the project

In order to buil this project you got to have Golang 1.24.2 (or some compatible version) installed.

1. Run the command `go build -o task-cli src/main.go'` to generate an executable file

## How to use the project

Here is an example of how to use this project:

```sh
# Adding a new task
./task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
./task-cli update 1 "Buy groceries and cook dinner"
./task-cli delete 1

# Marking a task as in progress or done
./task-cli mark-in-progress 1
./task-cli mark-done 1

# Listing all tasks
./task-cli list

# Listing tasks by status
./task-cli list done
./task-cli list todo
./task-cli list in-progress
```
