package main

import (
	"fmt"
	"os"

	"go-reloaded/internal/audit"
	"go-reloaded/internal/processor"
	"go-reloaded/pkg/fileio"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt> [--audit]")
		os.Exit(1)
	}

	processor := processor.New()

	if len(os.Args) > 3 && os.Args[3] == "--audit" {
		audit.New(processor).Run()
		return
	}

	content, err := fileio.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	result := processor.Process(content)

	if err := fileio.WriteFile(os.Args[2], result); err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}
}