package task

type Tasker interface {
	GetId(idx int) int
	GetName(idx int) string
	GetProgress(idx int) int
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

func (task *Task) Init(idtask int, nametask string, progressname int) {
	task.id = idtask
	task.name = nametask
	task.progress = progressname
}
