package model

type Task struct {
	Name     string `json:"name"`     // название задачи
	Progress int    `json:"progress"` // прогресс
}

type ListTask []*Task
