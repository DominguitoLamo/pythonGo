package main

import (
	"fmt"
	"os"

	"chonlam.com/pythonGo/code"
)

func main() {
	if (len(os.Args) <= 1) {
		fmt.Println("VM needs filename")
		return
	}

	path := os.Args[1]
	code.Parse(path)
}