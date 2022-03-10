package application

import (
	"log"

	"github.com/Giovanny472/gtask/internal/command"
	"github.com/Giovanny472/gtask/internal/config"
	"github.com/Giovanny472/gtask/internal/task"
	"github.com/Giovanny472/gtask/internal/ui"
	"github.com/Giovanny472/gtask/model"
)

// интерфейс для управления апп
type AppTasker interface {
	Config()
	Start()
}

// управление апп
type AppTask struct {

	// список commands. загрузка данных через json
	listCommand *model.ListCommand
	// список task. загрузка данных через json
	listTsk *model.ListTask

	// для настройки App
	configApp config.Configurator
	// управление task
	manTask task.ManagerTask
	// управление Parser command
	manCommand command.ManagerCommand
	// ui
	uiTerminal model.Shower
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
	aptask.listCommand = model.NewListCommand()

	// Manager Task
	aptask.manTask = task.NewManagerTask()

	// Manager Command
	aptask.manCommand = command.NewManagerCommand()

	// ui
	aptask.uiTerminal = ui.NewUITerminal()
}

func (aptask *AppTask) Config() {

	// загрузка commands апп
	err := aptask.configApp.Load(config.FileConfigName, &aptask.listCommand)
	if err != nil {
		log.Fatal(err)
	}

	// загрузка task
	err = apptask.configApp.Load(config.FileTasksName, &aptask.listTsk)
	if err != nil {
		log.Fatal(err)
	}

	// инициализация Manager Command
	aptask.manTask.Init(aptask.listTsk)

	// инициализация Manager Command
	aptask.manCommand.Init(aptask.listCommand, aptask.manTask)

	// загрузка комманды
	aptask.manCommand.Load()

	// Parse комманды
	aptask.manCommand.Parse()
}

func (aptask *AppTask) Start() {

	// запуск изменений
	aptask.manCommand.Execute()

	// сохранение изменений задач
	aptask.configApp.Save(config.FileTasksName, *aptask.manTask.GetListTasks())

	// ui show
	aptask.uiTerminal.Show(aptask.manTask.GetListTasks())
}
