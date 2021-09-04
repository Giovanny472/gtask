package task

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
	Update             = "u"
	Delete             = "d"
	Read               = "r"
	List               = "l"
)
