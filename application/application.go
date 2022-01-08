package application

import (
	"fmt"
	"log"

	"github.com/Giovanny472/gtask/model"
	"github.com/Giovanny472/gtask/pkg/config"
)

type AppTasker interface {
	Config()
	Start()
}

// управление апп
type AppTask struct {
	configApp   config.Configurator
	listCommand model.ListCommand
}

var apptask *AppTask

func New() AppTasker {

	if apptask == nil {

		// создание экземпляр AppTask
		apptask = &AppTask{}

		// инициализация
		apptask.init()
	}

	return apptask
}

func (aptask *AppTask) init() {

	// создание экземпляра настроек
	aptask.configApp = config.New()

	// создание экземпляра команд
	//aptask.listCommand = model.NewListCommand()
}

func (aptask *AppTask) Config() {

	// загрузка настроек апп
	err := aptask.configApp.Load(config.FileConfigName, &aptask.listCommand)
	if err != nil {
		log.Fatal(err)
	}

}

func (aptask *AppTask) Start() {

	fmt.Println(aptask.listCommand)
}
