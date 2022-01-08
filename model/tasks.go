package model

type Tasker interface {
	Name() string
	Progress() int
}

type ListTask []Tasker

type Task struct {
	name     string
	progress int
}

// ***** TASK *****
func New(idtask int, nametask string, progressname int) Tasker {

	var aTask Task

	aTask.name = nametask
	aTask.progress = progressname

	return &aTask
}

func (task *Task) Name() string {
	return task.name
}

func (task *Task) Progress() int {
	return task.progress
}
