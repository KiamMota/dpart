package main

import (
	"bufio"
	"fmt"
	"os"
	"dpart/core"
)

func main() {
	core.InterState.StartInternalState()	
	reader := bufio.NewReader(os.Stdin)
	for {
		print(">>> ")	
		input, _ := reader.ReadString('\n') 
		fmt.Printf("<<< %s\n", Dispatcher(&input))
	}
}
