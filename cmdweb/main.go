package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Giovanny472/gtask/internal/config"
)

type Task struct {
	Id       int
	Name     string
	User     string
	Progress int
}

var ListTask []*Task

func ReadJson() {

	ConfigApp := config.New()

	// загрузка task
	err := ConfigApp.Load(config.FileTasksWebName, &ListTask)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	// сайт
	http.HandleFunc("/gtask", gtaskfunc)

	// static файлы
	fs := http.FileServer(http.Dir("./web/static/"))
	http.Handle("/gtask/web/static/", http.StripPrefix("/gtask/web/static/", fs))

	log.Fatal("Error SiteWeb gtask -->", http.ListenAndServe(":8080", nil))
}

func gtaskfunc(w http.ResponseWriter, r *http.Request) {

	// чтение данных
	ReadJson()

	for idx, val := range ListTask {
		val.Id = idx + 1
	}

	// отправка и отображение данных в формате html
	strtmpl := "web/tmpl/gtask.html"
	template.Must(template.ParseFiles(strtmpl)).Execute(w, ListTask)
}
