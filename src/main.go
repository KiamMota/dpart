package main

import (
	"bufio"
	"dpart/commands"
	"dpart/commands/file"
	"fmt"
	"os"
	"strings"
)

type ProgramState struct {
	CurrentDirectory string
	UserDir          string
}

const (
	StrTrue  = "True"
	StrFalse = "False"
)

var UsedCommands []string

var progState ProgramState = ProgramState{}

func functionFileExists(param string) string {
	_, err := os.Stat(param)
	if err == nil {
		return StrTrue
	}
	return StrFalse
}

func dispatcher(input *string) string {
	args := strings.Fields(*input)
	if len(args) == 0 {
		return ""
	}

	command := args[0]

	switch command {
	case "exit":
		os.Exit(0)

	case "pwd":
		return progState.CurrentDirectory

	case "cd":
		if len(args) < 2 {
			return "uso: go <diretório>"
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	userdir, _ := os.UserHomeDir()
	pwd, _ := os.Getwd()
	progState.CurrentDirectory = pwd
	progState.UserDir = userdir

	for {
		fmt.Print("| > ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("erro:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		output := dispatcher(&input)
		if output != "" {
			fmt.Println(output)
		}
	}

}
