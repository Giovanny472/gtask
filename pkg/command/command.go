package command

import "flag"

// набор комманд

// Интерфейс
type Commander interface {
	Config()
	Parse()
}

// Реализация  интерфейса
type Command struct {
	Name  TypeCommand
	Value string
}

// Cписок комманд
type TypeCommand string

const (
	Create TypeCommand = "add"
	Update TypeCommand = "u"
	Delete TypeCommand = "d"
	Read   TypeCommand = "r"
	List   TypeCommand = "l"
)

func (comm *Command) Config() {
	flag.StringVar(&comm.Value, string(comm.Name), "", "")
}

func (comm *Command) Parse() {

}
