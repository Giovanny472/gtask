package application

import (
	"github.com/Giovanny472/gtask/pkg/config"
)

type AppTasker interface {
	Config()
	Start()
}

// управление апп
type AppTask struct {
	configApp config.Configurator
}

func (appTask *AppTask) init() {

	// создание экземпляр настроек
	appTask.configApp = config.New()
}

func (appTask *AppTask) Config() {

	// загрузка настроек апп
	appTask.configApp.Load(config.FileName, OBJECT_JSON)
}

func (appTask *AppTask) Start() {

}

func New() AppTasker {

	// создание экземпляр AppTask
	appTask := &AppTask{}
	appTask.init()

	return appTask
}
