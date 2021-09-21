package app

type Application interface {
	Start()
}

type AppTask struct {
}

func (app *AppTask) Start() {

}
