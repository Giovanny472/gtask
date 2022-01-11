package task

import (
	"strings"

	"github.com/Giovanny472/gtask/model"
)

type ManagerTask interface {

	// *****CRUD для задач*****

	// создание задачи
	Create(value string)
	// чтение задачи
	Read(value string)
	// обновление задачи
	Update()
	// удаление задачи
	Delete()
}

type managerTsk struct {
	listTask model.ListTask
}

func NewManagerTask() ManagerTask {
	return &managerTsk{}
}

// ***** MANAGER *****

// создание задачи
//func (mang *managerTsk) Create(task model.Task) {
func (mang *managerTsk) Create(value string) {

	//mang.listTask = append(mang.listTask, &task)
}

// чтение задачи
//func (mang *managerTsk) Read(value string) {
func (mang *managerTsk) Read(value string) {
	/*
		// получаем
		// id slice,
		// task,
		// bool - сущесвтует ли задача
		_, atask, ok := mang.isTask(value)
		if !ok {
			fmt.Println("нет задачи.......... 0")
			return
		}

		fmt.Println(atask.Name, "..........", atask.Progress)
	*/
}

// обновление задачи
//func (mang *managerTsk) Update(value string, task model.Task) bool {
func (mang *managerTsk) Update() {
	/*

		if &task == nil {
			return false
		}

		idx, atask, ok := mang.isTask(value)
		if !ok {
			return false
		}

		// обновление
		mang.listTask[idx] = &atask

		return true
	*/
}

// удаление задачи
//func (mang *managerTsk) Delete(value string) bool {
func (mang *managerTsk) Delete() {
	/*
		idx, _, ok := mang.isTask(value)
		if !ok {
			return false
		}

		mang.remove(idx)
		return true
	*/
}

func (mang *managerTsk) isTask(value string) (int, model.Task, bool) {

	if len(mang.listTask) == 0 {
		return -1, model.Task{}, false
	}

	for idx, atsk := range mang.listTask {

		ok := strings.Contains(atsk.Name, value)
		if ok {
			return idx, *atsk, true
		}
	}

	return -1, model.Task{}, false
}

func (mang *managerTsk) remove(idx int) {

	newLisTask := append(mang.listTask[:idx], mang.listTask[idx+1:]...)
	mang.listTask = newLisTask
}
