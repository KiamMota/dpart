package main

import (
	"bufio"
	"dpart/cli"
	"dpart/core"
	"fmt"
	"os"
)

func main() {
	core.InterState.StartInternalState()	
	reader := bufio.NewReader(os.Stdin)
	for {
		print(">>> ")	
		input, _ := reader.ReadString('\n') 
		fmt.Printf("<<< %s\n", cli.Dispatcher(&input))
	}
}
