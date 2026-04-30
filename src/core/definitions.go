package core

const (
	StrTrue = "True"
	StrFalse = "False"
	StrNil = "Nil"
)

func MacroTable(macro string) string {
	switch macro {
	case "pwd":
		return InterState.CurrentDirectory
	case "home":
		return InterState.UserHomeDir
	default: return macro
	}
}
