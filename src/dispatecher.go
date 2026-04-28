package main

import (
	"dpart/commands"
	"dpart/commands/file"
	"dpart/core"
	"os"
	"strings"
)

func Dispatcher(input *string ) string {
	args := strings.Fields(*input)
	if len(args) == 0 {
		return ""
	}

	command := args[0]

	switch command {
	case "exit":
		os.Exit(0)

	case "pwd":
		return core.InterState.CurrentDirectory

	case "cd":
		if len(args) < 2 {
			return "cd <dir>"
		}
		return commands.ChangeCurrentDirectory(args[1])
	case "write":
		if len(args) < 2 {
			return "too few args"
		}
		return file.FileWrite(args[1])
	case "read":
		return file.FileRead(args[1])
	case "exists":
		return commands.Exists(args[1])
	case "inspect":
		return commands.Inspect(args[1])
	}

	return "comando desconhecido"
}



