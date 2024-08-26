https://roadmap.sh/projects/task-tracker

add: go run main.go add -name={"something"}

list: go run main.go list

list tasks with status "done": go run main.go list-done

list tasks with status "not done": go run main.go list-not-done

list tasks with status "in progress": go run main.go list-in-progress

update tasks: go run main.go update-status {id} {status}

delete task: go run main.go delete {id}
