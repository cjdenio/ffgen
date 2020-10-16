package main

import (
	"fmt"
	"os"

	"github.com/cjdenio/ffgen/pkg/help"
)

func main() {
	// If there are fewer than 2 arguments, show help
	if len(os.Args) < 3 {
		help.PrintHelp()
		return
	}

	directory := os.Args[1]
	file := os.Args[2]

	fmt.Println(directory)
	fmt.Println(file)
}
