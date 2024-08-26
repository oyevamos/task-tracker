package cmd

import (
	"fmt"
	"time"
)

func AddTask(taskName string) {
	tasks := loadTask()

	newID := len(tasks) + 1
	currentTime := time.Now()
	newTask := Task{
		ID:        newID,
		Name:      taskName,
		Status:    "not done",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	tasks = append(tasks, newTask)
	saveTasks(tasks)

	fmt.Println("Task added succeffully")
}
