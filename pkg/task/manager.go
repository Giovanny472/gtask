package task

import (
	"strings"

	"github.com/Giovanny472/gtask/model"
)

type ManagerTask interface {

	// *****CRUD для задач*****

	// создание задачи
	Create(task model.Tasker)
	// чтение задачи
	Read(value string) model.Tasker
	// обновление задачи
	Update(value string, task model.Tasker) bool
	// удаление задачи
	Delete(value string) bool
}

type Manager struct {
	listTask model.ListTask
}

// ***** MANAGER *****

// создание задачи
func (mang *Manager) Create(task model.Tasker) {

	mang.listTask = append(mang.listTask, task)
}

// чтение задачи
func (mang *Manager) Read(value string) model.Tasker {

	// получаем
	// id slice,
	// task,
	// bool - сущесвтует ли задача
	_, atask, ok := mang.isTask(value)
	if !ok {
		return nil
	}
	return atask
}

// обновление задачи
func (mang *Manager) Update(value string, task model.Tasker) bool {

	if task == nil {
		return false
	}

	idx, atask, ok := mang.isTask(value)
	if !ok {
		return false
	}

	// обновление
	mang.listTask[idx] = atask

	return true
}

// удаление задачи
func (mang *Manager) Delete(value string) bool {

	idx, _, ok := mang.isTask(value)
	if !ok {
		return false
	}

	mang.remove(idx)
	return true
}

func (mang *Manager) isTask(value string) (int, model.Tasker, bool) {

	if len(mang.listTask) == 0 {
		return -1, nil, false
	}

	for idx, atsk := range mang.listTask {

		ok := strings.Contains(atsk.Name(), value)
		if ok {
			return idx, atsk, true
		}
	}

	return -1, nil, false
}

func (mang *Manager) remove(idx int) {

	newLisTask := append(mang.listTask[:idx], mang.listTask[idx+1:]...)
	mang.listTask = newLisTask
}
