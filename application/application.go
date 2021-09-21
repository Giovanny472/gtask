package application

import "fmt"

// управление апп
type AppTask struct {
}

func (appTask *AppTask) Start() {
	fmt.Println("Start app")
}
