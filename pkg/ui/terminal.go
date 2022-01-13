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

	for _, value := range *listask {
		fmt.Println(value)
	}
}
