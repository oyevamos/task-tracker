package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadTask() []Task {
	var tasks []Task

	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		return tasks
	}

	file, err := os.ReadFile(taskFile)
	if err != nil {
		fmt.Println("Error reading tasks file:", err)
		return tasks
	}

	err = json.Unmarshal(file, &tasks)
	if err != nil {
		fmt.Println("Error parsing file")
	}

	return tasks
}

func saveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error serializing tasks")
		return
	}

	err = os.WriteFile(taskFile, data, 0644)
	if err != nil {
		fmt.Println("Error writing tasks file")
	}
}
