package main

import (
	"fmt"
	"os"

	"chonlam.com/pythonGo/codeStream"
)

func main() {
	if (len(os.Args) <= 1) {
		fmt.Println("VM needs filename")
		return
	}

	path := os.Args[1]
	buffer := codeStream.CreateStringBuffer(path)
	fmt.Printf("magic number is %x\n", buffer.ReadInt())
}