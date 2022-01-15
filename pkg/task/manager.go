package task

import (
	"strconv"
	"strings"

	"github.com/Giovanny472/gtask/model"
)

type ManagerTask interface {

	// *****CRUD для задач*****

	// инициализация
	Init(lstTsk *model.ListTask)

	// создание задачи
	Create(value string, params []string)
	// чтение задачи
	Read(value string)
	// обновление задачи
	Update(value string)
	// удаление задачи
	Delete(value string)

	// список задач
	GetListTasks() *model.ListTask
}

type managerTsk struct {
	listTask *model.ListTask
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
func (mang *managerTsk) Create(value string, params []string) {

	var atsk model.Task

	// получаем новые значения
	// [1] название задачи
	atsk.Name = value
	// [2] процент
	if len(params) > 0 {
		aprog, err := strconv.Atoi(params[0])
		if err == nil {
			atsk.Progress = aprog
		}
	}

	// проверка задачи
	_, ok := mang.isTask(atsk.Name)
	if !ok {

		// новый список
		//var listChangedTasks model.ListTask
		listChangedTasks := append(*mang.listTask, &atsk)

		// обновление списка задач
		mang.listTask = &listChangedTasks
	}

}

// чтение задачи
func (mang *managerTsk) Read(value string) {

	if len(value) == 0 {
		return
	}

	if value == "all" {
		return
	}
}

// обновление задачи
func (mang *managerTsk) Update(value string) {

	alistparams := strings.Split(value, ",")
	if len(alistparams) < 3 {
		return
	}

	// получаем название задачи
	aName := alistparams[0]

	// получаем новые значения
	aNamNew := alistparams[1]
	aProNew, isOkProg := strconv.Atoi(alistparams[2])

	// найдем задачу из списка
	atask, ok := mang.isTask(aName)
	if !ok {
		return
	}

	// обновление
	atask.Name = aNamNew

	if isOkProg == nil {
		atask.Progress = aProNew
	}

}

// удаление задачи
func (mang *managerTsk) Delete(value string) {

	atask, ok := mang.isTask(value)
	if !ok {
		return
	}

	mang.remove(atask)
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
func (mang *managerTsk) remove(atsk *model.Task) {

	// новый список
	var listChangedTasks model.ListTask
	for _, avalTask := range *mang.listTask {

		if avalTask == atsk {
			continue
		}
		listChangedTasks = append(listChangedTasks, avalTask)
	}
	mang.listTask = &listChangedTasks

}

func (mang *managerTsk) GetListTasks() *model.ListTask {
	return mang.listTask
}
