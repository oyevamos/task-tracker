package cmd

import (
	"fmt"
	"time"
)

func UpdateTaskStatus(id int, status string) {
	tasks := loadTask()
	updated := false
	currentTime := time.Now()

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = currentTime
			updated = true
			break
		}
	}

	if !updated {
		fmt.Printf("Task %d not found", id)
		return
	}

	saveTasks(tasks)
	fmt.Printf("Task %d status updated to %s\n", id, status)
}
