package main

import "github.com/Giovanny472/gtask/application"

func main() {

	// создание апп
	aTaskApp := application.NewApp()

	// настройка апп
	aTaskApp.Config()

	// старт апп
	aTaskApp.Start()

}
