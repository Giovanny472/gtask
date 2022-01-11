package command

import (
	"flag"

	"github.com/Giovanny472/gtask/model"
	"github.com/Giovanny472/gtask/pkg/task"
)

// управление коммандами
type ManagerCommand interface {
	Load(listCom *model.ListCommand)
	Parse()
	Execute(mantask task.ManagerTask, listCom *model.ListCommand)
}

// структура для управления коммандами
type managerCom struct {
}

// создание экземпляра
func NewManagerCommand() ManagerCommand {
	return &managerCom{}
}

// загрузка flags
func (mang *managerCom) Load(listCom *model.ListCommand) {

	for _, item := range *listCom {

		// настройка Flag
		flag.StringVar(&item.ValueUsr, item.Instruction, item.ValueDefault, item.Description)
	}
}

// Par flags
func (mang *managerCom) Parse() {
	flag.Parse()
}

func (mang *managerCom) Execute(mantask task.ManagerTask, listCom *model.ListCommand) {

	for _, item := range *listCom {

		if item.ValueUsr != "" {
			mang.executeProcess(item.Name, item.ValueUsr, mantask)
		}

	}
}

func (mang *managerCom) executeProcess(nametask string, valuetask string, mantask task.ManagerTask) {

	// процесс
	switch nametask {

	// создание
	case model.CommandAdd:
		mantask.Create(valuetask)

	// обновление
	case model.CommandUpd:

	case model.CommandDel:

	case model.CommandRea:

	case model.CommandPro:

	}

}
