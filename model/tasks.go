package model

type Task struct {
	Name     string `json:"name"`     // название задачи
	Progress string `json:"progress"` // прогресс
}

type ListTask []*Task
