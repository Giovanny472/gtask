package model

type Color string

const (
	Black     Color = "\033[0;30m"
	Red       Color = "\033[0;31m"
	RedBold   Color = "\033[1;31m"
	Green     Color = "\033[0;32m"
	GreenBold Color = "\033[1;32m"
	Yellow    Color = "\033[0;33m"
	Blue      Color = "\033[0;34m"
	Purple    Color = "\033[0;35m"
	Cyan      Color = "\033[0;36m"
	White     Color = "\033[0;37m"
)
