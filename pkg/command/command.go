package command

// набор комманд

// Интерфейс
type Commander interface {
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
