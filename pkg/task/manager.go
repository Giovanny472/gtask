package task

import "github.com/Giovanny472/gtask/model"

type ManagerTask interface {
	Create()
	Read()
	Update()
	Delete()
}

type Manager struct {
	listTask model.ListTask
}

func (mang *Manager) Create() {
}

func (mang *Manager) Read() {
}

func (mang *Manager) Update() {
}

func (mang *Manager) Delete() {
}
