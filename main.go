package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"task-tracker/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'add', 'list', or 'update-status' subcommands.")
		return
	}

	switch os.Args[1] {
	case "add":
		flagSet := flag.NewFlagSet("add", flag.ExitOnError)
		taskName := flagSet.String("name", "", "Name of the task")
		flagSet.Parse(os.Args[2:])
		if *taskName == "" {
			fmt.Println("Task name is requires")
			return
		}
		cmd.AddTask(*taskName)

	case "list":
		cmd.ListTask()
	case "list-done":
		cmd.ListDoneTasks()
	case "list-not-done":
		cmd.ListNotDoneTasks()
	case "list-in-progress":
		cmd.ListInProgressTasks()
	case "update-status":
		if len(os.Args) < 4 {
			fmt.Println("Usage: update-status <task ID> <status>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		status := os.Args[3]
		if status != "done" && status != "in progress" && status != "not done" {
			fmt.Println("Invalid status. Valid statuses are: 'done', 'in progress', 'not done'")
			return
		}
		cmd.UpdateTaskStatus(id, status)
	case "delete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		cmd.DeleteTask(id)
	default:
		fmt.Println("Expected 'add', 'list', 'list-done', 'list-not-done', 'list-in-progress', or 'update-status' subcommands.")
	}
}
