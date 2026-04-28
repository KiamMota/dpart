package core

import (
	"os"
	"strings"
)


func DispatcherFunc(command string) string {
	args := strings.Split(command, " ")	
	command = args[0]
	switch command { 
		case "exit":
			os.Exit(0)
	}
	return ""

}
