package command

import (
	"flag"

	"github.com/Giovanny472/gtask/internal/task"
	"github.com/Giovanny472/gtask/model"
)

// управление коммандами
type ManagerCommand interface {
	Init(lstCom *model.ListCommand,
		manTsk task.ManagerTask)
	Load()
	Parse()
	Execute()
}

// структура для управления коммандами
type managerCom struct {
	listCom *model.ListCommand
	manTask task.ManagerTask
}

// создание экземпляра
func NewManagerCommand() ManagerCommand {
	return &managerCom{}
}

func (mang *managerCom) Init(lstCom *model.ListCommand, manTsk task.ManagerTask) {

	mang.listCom = lstCom
	mang.manTask = manTsk
}

// загрузка flags
func (mang *managerCom) Load() {

	for _, item := range *mang.listCom {

		// настройка Flag
		flag.StringVar(&item.ValueUsr, item.Instruction, item.ValueDefault, item.Description)
	}
}

// Par flags
func (mang *managerCom) Parse() {

	flag.Parse()
}

func (mang *managerCom) Execute() {

	for _, item := range *mang.listCom {

		if item.ValueUsr != "" {
			mang.executeProcess(item.Name, item.ValueUsr)
		}
	}
}

func (mang *managerCom) executeProcess(nametask string, valuetask string) {

	// процесс
	switch nametask {

	// создание
	case model.CommandAdd:
		mang.manTask.Create(valuetask, flag.Args())

	// обновление
	case model.CommandUpd:
		mang.manTask.Update(valuetask, flag.Args())

	case model.CommandDel:
		mang.manTask.Delete(valuetask)

	case model.CommandRea:
		mang.manTask.Read(valuetask)
	}

}
