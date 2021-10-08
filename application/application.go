package application

import (
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

func (aptask *AppTask) init() {

	// создание экземпляр настроек
	aptask.configApp = config.New()
}

func (aptask *AppTask) Config() {

	// загрузка настроек апп
	err := aptask.configApp.Load(config.FileName, &aptask.listCommand)
	if err != nil {
		log.Fatal(err)
	}

}

func (aptask *AppTask) Start() {

}

func New() AppTasker {

	if apptask == nil {

		// создание экземпляр AppTask
		apptask = &AppTask{}

		// инициализация
		apptask.init()
	}

	return apptask
}
