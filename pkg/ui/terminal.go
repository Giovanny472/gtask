package ui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/Giovanny472/gtask/model"
)

type uiterminal struct {
}

func NewUITerminal() model.Shower {
	return &uiterminal{}
}

func (uiterm *uiterminal) Show(listask *model.ListTask) {

	uiterm.clearTerminal()

	fmt.Println(model.Cyan, "********************************")
	fmt.Println(model.Cyan, "*      З А Д А Ч И             *")
	fmt.Println(model.Cyan, "********************************")

	//   список задачи
	var acolor model.Color
	for idx, value := range *listask {

		if value.Progress > 90 {
			acolor = model.GreenBold
		} else if value.Progress < 50 {
			acolor = model.RedBold
		} else {
			acolor = model.Cyan
		}

		// счетчик
		fmt.Print(model.Cyan, "[", model.Cyan, idx+1, "]")

		if idx < 10 {
			fmt.Print("  ")
		} else {
			fmt.Print(" ")
		}

		// название задачи
		fmt.Print(model.Cyan, value.Name)
		fmt.Print(model.Cyan, " ..... ")

		fmt.Println(acolor, value.Progress, "%")
		fmt.Print(model.ColorReset)
	}
}

func (uiterm *uiterminal) clearTerminal() {

	var cmd *exec.Cmd

	if runtime.GOOS == "linux" {
		cmd = exec.Command("clear")
	} else if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
