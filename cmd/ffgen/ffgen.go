package main

import (
	"fmt"
	"os"

	"path/filepath"

	"github.com/cjdenio/ffgen/pkg/fontfile"
	"github.com/cjdenio/ffgen/pkg/gen"
	"github.com/cjdenio/ffgen/pkg/help"
	"github.com/cjdenio/ffgen/pkg/write"
)

func main() {
	// If there are fewer than 2 arguments, show help
	if len(os.Args) < 3 {
		help.PrintHelp()
		return
	}

	// Convert arguments to absolute paths
	directory, _ := filepath.Abs(os.Args[1])
	file, _ := filepath.Abs(os.Args[2])

	// Get path of CSS file
	path := filepath.Dir(file)

	// Scan directory for font files
	files, err := fontfile.SearchDirectory(directory)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("ℹ️ Found %d font files\n", len(files))

	rules := gen.GenerateRules(files, path)

	err = write.WriteToFile(file, rules)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("🎉 Successfully created @font-face rules!")
}
