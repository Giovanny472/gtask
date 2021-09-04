package task

type Tasker interface {
	GetId() int
	GetName() string
	GetProgress() int
}

type Task struct {
	id       int
	name     string
	progress int
}

type ListTask []*Task

func (task *Task) GetId() int {
	return task.id
}

func (task *Task) GetName() string {
	return task.name
}

func (task *Task) GetProgress() int {
	return task.progress
}
