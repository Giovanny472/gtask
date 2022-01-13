package ui

import (
	"fmt"

	"github.com/Giovanny472/gtask/model"
)

type uiterminal struct {
}

func NewUITerminal() model.Shower {
	return &uiterminal{}
}

func (uiterm *uiterminal) Show(listask *model.ListTask) {

	//fmt.Println(model.Cyan, "********************************")
	//fmt.Println(model.Cyan, "*      З А Д А Ч И             *")
	//fmt.Println(model.Cyan, "********************************")
	fmt.Println("********************************")
	fmt.Println("*      З А Д А Ч И             *")
	fmt.Println("********************************")

	//   список задачи
	//var acolor model.Color
	for idx, value := range *listask {

		//if value.Progress > 90 {
		//	acolor = model.GreenBold
		//} else if value.Progress < 50 {
		//	acolor = model.RedBold
		//} else {
		//	acolor = model.Cyan
		//}

		//fmt.Print(model.Cyan, "[", model.Cyan, idx+1, "]")
		fmt.Print("[", idx+1, "]")

		if idx < 10 {
			fmt.Print("  ")
		} else {
			fmt.Print(" ")
		}

		// название задачи
		//fmt.Print(model.Cyan, value.Name)
		fmt.Print(value.Name)
		fmt.Print(" ..... ")

		//fmt.Println(acolor, value.Progress, "%")
		fmt.Println(value.Progress, "%")
	}
}
