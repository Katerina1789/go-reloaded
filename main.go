package main

import (
	"fmt"    // For printing messages
	"os"     // For file operations and command line args
	"strings" // For text manipulation

	"go-reloaded/internal" // Our FSM and transformation logic
)

func main() {
	// Check if user provided required input and output file arguments
	if len(os.Args) < 3 {
		fmt.Println("Error: Missing required arguments")
		fmt.Println("Usage: go run . <input.txt> <output.txt> [--audit]")
		os.Exit(1) // Exit with error code
	}

	// Check if user wants audit mode (testing)
	if len(os.Args) > 3 && os.Args[3] == "--audit" {
		runAuditMode() // Run built-in tests
		return        // Exit after audit
	}

	// Read the input file content
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1) // Exit if file can't be read
	}

	// Process the text through our FSM
	result := processText(string(content))

	// Write the transformed text to output file
	err = os.WriteFile(os.Args[2], []byte(result), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1) // Exit if file can't be written
	}
}

func processText(text string) string {
	fsm := internal.NewFSM()    // Create new FSM instance
	words := tokenize(text)     // Split text into tokens

	// Feed each word to the FSM for processing
	for _, word := range words {
		fsm.ProcessWord(word) // FSM applies rules based on current state
	}

	return fsm.GetOutput() // Get final processed text
}

func tokenize(text string) []string {
	// Split text by spaces to get individual words
	words := strings.Fields(text)
	result := make([]string, 0) // Create empty slice for results
	
	// Process each word, combining range patterns
	for i := 0; i < len(words); i++ {
		word := words[i]
		// Check if this word starts a range pattern like "(up," "(low," "(cap,"
		if (word == "(up," || word == "(low," || word == "(cap,") && i+1 < len(words) {
			// Combine with next word to form complete pattern like "(up, 2)"
			combined := word + " " + words[i+1]
			result = append(result, combined) // Add combined pattern
			i++                              // Skip next word since we combined it
		} else {
			result = append(result, word) // Add word as-is
		}
	}
	
	return result // Return tokenized words
}

func runAuditMode() {
	fmt.Println("Running audit mode...")
	
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Hex conversion", "1E (hex) files were added", "30 files were added"},
		{"Binary conversion", "It has been 10 (bin) years", "It has been 2 years"},
		{"Uppercase", "Ready, set, go (up) !", "Ready, set, GO!"},
		{"Lowercase", "I should stop SHOUTING (low)", "I should stop shouting"},
		{"Capitalize", "Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"Range uppercase", "This is so exciting (up, 2)", "This is SO EXCITING"},
		{"Article correction", "There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"Punctuation spacing", "I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"Grouped punctuation", "I was thinking ... You were right", "I was thinking... You were right"},
		{"Quotes single word", "I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"Quotes multiple words", "As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
		{"Complex example", "it (cap) was the best of times, it was the worst of times (up)", "It was the best of times, it was the worst of TIMES"},
	}
	
	passed := 0
	for _, tc := range testCases {
		result := processText(tc.input)
		if result == tc.expected {
			fmt.Printf("✅ %s: PASS\n", tc.name)
			passed++
		} else {
			fmt.Printf("❌ %s: FAIL\n", tc.name)
			fmt.Printf("   Expected: %s\n", tc.expected)
			fmt.Printf("   Got:      %s\n", result)
		}
	}
	
	fmt.Printf("\nAudit Results: %d/%d tests passed\n", passed, len(testCases))
	if passed == len(testCases) {
		fmt.Println("🎉 All tests passed!")
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}