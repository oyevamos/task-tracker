package cmd

import "fmt"

func DeleteTask(id int) {
	tasks := loadTask()
	indexToDelete := -1

	for i, task := range tasks {
		if task.ID == id {
			indexToDelete = i
			break
		}
	}
	if indexToDelete == -1 {
		fmt.Printf("Task with ID %d not found", id)
		return
	}
	tasks = append(tasks[:indexToDelete], tasks[indexToDelete+1:]...)

	saveTasks(tasks)
	fmt.Printf("Task with ID %d has been deleted\n", id)
}
