package model

import (
	"fmt"
)

// набор комманд

// Интерфейс
type Commander interface {
	Config()
	Parse()
}

// Реализация  интерфейса
type Command struct {
	Name         string `json:"name"`         // название команды
	Instruction  string `json:"instruction"`  // символ команды
	ValueDefault string `json:"valuedefault"` // значение команды по умолчанию
	ValueUsr     string `json:"valueUsr"`     // значение, которое дает пользователю
	Description  string `json:"description"`  // описание команды
}

type ListCommand []Command

// Cписок комманд
type TypeCommand string

const (
	Create TypeCommand = "a"
	Update TypeCommand = "u"
	Delete TypeCommand = "d"
	Read   TypeCommand = "r"
	List   TypeCommand = "l"
)

func NewListCommand() *ListCommand {
	return &ListCommand{}
}

func (listcom ListCommand) Config() {

	for comm := range listcom {

		fmt.Printf("comm: %v\n", comm)
		//flag.StringVar(&comm.ValueUsr, comm.Instruction, comm.ValueDefault, comm.Description)

	}

}

func (comm *Command) Parse() {

}
