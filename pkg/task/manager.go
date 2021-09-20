package task

type ManagerTask interface {
	Create()
	Read()
	Update()
	Delete()
}

type Manager struct {
	listTask ListTask
}

func (mang *Manager) Create() {
}

func (mang *Manager) Read() {
}

func (mang *Manager) Update() {
}

func (mang *Manager) Delete() {
}
