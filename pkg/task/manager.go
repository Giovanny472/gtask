package task

import (
	"fmt"
	"strings"

	"github.com/Giovanny472/gtask/model"
)

type ManagerTask interface {

	// *****CRUD для задач*****

	// инициализация
	Init(lstTsk *model.ListTask)

	// создание задачи
	Create(value string)
	// чтение задачи
	Read()
	// обновление задачи
	Update(value string)
	// удаление задачи
	Delete(value string)

	// список задач
	GetListTasks() *model.ListTask
}

type managerTsk struct {
	listTask         *model.ListTask
	listChangedTasks model.ListTask
}

func NewManagerTask() ManagerTask {
	return &managerTsk{}
}

// ***** MANAGER *****

// инициализация
func (mang *managerTsk) Init(lstTsk *model.ListTask) {

	// список tasks
	mang.listTask = lstTsk
}

// создание задачи
func (mang *managerTsk) Create(value string) {

	var atsk model.Task
	atsk.Name = "demo03"
	atsk.Progress = 80

	//mang.listTask
	mang.listChangedTasks = append(*mang.listTask, &atsk)

	// обновление списка задач
	mang.listTask = &mang.listChangedTasks

	fmt.Println("CREATE")
}

// чтение задачи
func (mang *managerTsk) Read() {
	// без изменения данных
	// поэтому без реализация метода
}

// обновление задачи
func (mang *managerTsk) Update(value string) {

	atask, ok := mang.isTask(value)
	if !ok {
		return
	}

	// обновление

	atask.Name = "nonono"
	atask.Progress = 101010

	fmt.Println("UPDATE")
}

// удаление задачи
func (mang *managerTsk) Delete(value string) {
	/*
		idx, _, ok := mang.isTask(value)
		if !ok {
			return false
		}

		mang.remove(idx)
		return true
	*/
}

// существует ли задача
func (mang *managerTsk) isTask(value string) (*model.Task, bool) {

	if len(*mang.listTask) == 0 {
		return &model.Task{}, false
	}

	for _, atsk := range *mang.listTask {

		ok := strings.Contains(atsk.Name, value)
		if ok {
			return atsk, true
		}
	}

	return &model.Task{}, false
}

// удаление задачи
func (mang *managerTsk) remove(idx int) {
	/*
		newLisTask := append(mang.listTask[:idx], mang.listTask[idx+1:]...)
		mang.listTask = newLisTask
	*/
}

func (mang *managerTsk) GetListTasks() *model.ListTask {
	return mang.listTask
}
