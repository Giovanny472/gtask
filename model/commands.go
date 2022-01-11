package model

const (
	CommandAdd string = "add"
	CommandUpd string = "update"
	CommandDel string = "delete"
	CommandRea string = "read"
	CommandPro string = "progress"
)

// Реализация  интерфейса
type Command struct {
	Name         string `json:"name"`         // название команды
	Instruction  string `json:"instruction"`  // символ команды
	ValueDefault string `json:"valuedefault"` // значение команды по умолчанию
	ValueUsr     string `json:"valueUsr"`     // значение, которое дает пользователю
	Description  string `json:"description"`  // описание команды
}

type ListCommand []*Command

func NewListCommand() *ListCommand {
	return new(ListCommand)
}
