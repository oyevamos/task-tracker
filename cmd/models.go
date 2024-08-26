package cmd

import "time"

type Task struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at""`
	UpdatedAt time.Time `json:"updated-at"`
}

var taskFile = "task.json"
