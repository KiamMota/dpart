package cli 

import (
	"dpart/cli/commands"
	"dpart/core"
	"fmt"
	"os"
	"strings"
)

func parseMacro(macro string) string {
	if macro [0] != '$'{ return macro}
		if (macro [1] == '(' && macro[len(macro) -1] == ')'){
		pos := strings.Index(macro, ")")
		if pos == -1 {
			return macro
		}
		macroWord := macro[2:pos]
		fmt.Println("expanding macro: ", macroWord)
		return core.MacroTable(macroWord)	
	}
	return macro
}

func Dispatcher(input *string ) string {
	args := strings.Fields(*input)
	if len(args) == 0 {
		return ""
	}
	if args[0]== "lua" {
		return "lua is not implemented yet!"
	}
	command := args[0]

	for i := 1; i < len(args); i++{
		args[i] =parseMacro(args[i])	
	}	

	switch command {
	case "clear":
		return ""
	case "exit":	
		os.Exit(0)
	case "pwd":
		return core.InterState.CurrentDirectory
	case "get":
		return commands.Get(args[1])
	case "put":
		if len(args) == 2{
			return commands.Put(args[1], "")
		}
	}

	return "comando desconhecido"
}
 


