package main

import (
	"bufio"
	"dpart/core"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	userdir, _ := os.UserHomeDir()
	pwd, _ := os.Getwd()
	core.InterState.CurrentDirectory = pwd
	core.InterState.UserDir = userdir

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
		output := Dispatcher(&input)
		if output != "" {
			fmt.Println(output)
		}
	}

}
