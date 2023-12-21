package main

import (
	"fmt"
	"os"

	"chonlam.com/pythonGo/code"
	"chonlam.com/pythonGo/runtime"
)

func main() {
	if (len(os.Args) <= 1) {
		fmt.Println("VM needs filename")
		return
	}

	path := os.Args[1]
	codeObject := code.Parse(path)
	interpreter := new(runtime.Interpreter)
	interpreter.Run(codeObject)
}