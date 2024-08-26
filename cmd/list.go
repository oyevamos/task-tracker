package cmd

import (
	"fmt"
)

func ListTask() {
	tasks := loadTask()
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Name: %s, Status: %s\n", task.ID, task.Name, task.Status)
	}
}

func ListDoneTasks() {
	tasks := loadTask()
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("Done Tasks:")
	for _, task := range tasks {
		if task.Status == "done" {
			fmt.Printf("ID: %d, Name: %s\n", task.ID, task.Name)
		}
	}
}

func ListNotDoneTasks() {
	tasks := loadTask()
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	fmt.Println("Not Done Tasks:")
	for _, task := range tasks {
		if task.Status == "not done" {
			fmt.Printf("ID: %d, Name: %s\n", task.ID, task.Name)
		}
	}
}

func ListInProgressTasks() {
	tasks := loadTask()
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("In Progress Tasks:")
	for _, task := range tasks {
		if task.Status == "in progress" {
			fmt.Printf("ID, %d, Name %s\n", task.ID, task.Name)
		}
	}
}
