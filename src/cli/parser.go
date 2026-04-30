package cli

import (
	"dpart/cli/commands"
	"dpart/core"
	"os"
	"regexp"
	"strings"
)

var macroRegex = regexp.MustCompile(`\$\((.*?)\)`)

func expandMacros(input string) string {
	return macroRegex.ReplaceAllStringFunc(input, func(match string) string {
		key := strings.TrimSpace(match[2 : len(match)-1])
		return core.MacroTable(key)
	})
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
		args[i] = expandMacros(args[i])	
	}	

	switch command {
	case "clear":
		return ""
	case "exit":	
		os.Exit(0)
	case "cd":
		return commands.Cd(args[1])
	case "pwd":
		return core.InterState.CurrentDirectory
	case "get":
		return commands.Get(args[1])
	case "put":
		if len(args) == 2{
			return commands.Put(args[1], "")
		}
	case "remove":
		return commands.Remove(args[1])
	}

	return "comando desconhecido"
}
 


