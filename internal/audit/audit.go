package audit

import (
	"fmt"
	"os"
	"strings"
)

type TestCase struct {
	Name     string
	Input    string
	Expected string
}

type Auditor struct {
	processor interface{ Process(string) string }
}

func New(processor interface{ Process(string) string }) *Auditor {
	return &Auditor{processor: processor}
}

func (a *Auditor) Run() {
	tests := a.getTestCases()
	passed := 0
	
	fmt.Printf("Running audit mode...\nTotal tests: %d\n\n", len(tests))
	
	for i, test := range tests {
		fmt.Printf("Test %d: %s\n", i+1, test.Name)
		
		result := a.processor.Process(test.Input)
		if strings.TrimSpace(result) == strings.TrimSpace(test.Expected) {
			fmt.Printf("  PASS\n\n")
			passed++
		} else {
			fmt.Printf("  FAIL\n")
			fmt.Printf("  Expected: %q\n", test.Expected)
			fmt.Printf("  Got:      %q\n\n", result)
		}
	}
	
	fmt.Printf("Results: %d/%d tests passed\n", passed, len(tests))
	if passed != len(tests) {
		fmt.Printf("%d tests failed\n", len(tests)-passed)
		os.Exit(1)
	}
	fmt.Println("All tests passed!")
}

func (a *Auditor) getTestCases() []TestCase {
	return []TestCase{
		{
			Name:     "Mixed Casing Rules",
			Input:    "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			Expected: "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			Name:     "Number Conversions",
			Input:    "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure.",
			Expected: "I have to pack 5 outfits. Packed 26 just to be sure.",
		},
		{
			Name:     "Punctuation Spacing",
			Input:    "Don not be sad ,because sad backwards is das . And das not good",
			Expected: "Don not be sad, because sad backwards is das. And das not good",
		},
		{
			Name:     "Quotes and Articles",
			Input:    "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			Expected: "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
		{
			Name:     "Basic hex conversion",
			Input:    "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			Expected: "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			Name:     "Article correction",
			Input:    "There is no greater agony than bearing a untold story inside you.",
			Expected: "There is no greater agony than bearing an untold story inside you.",
		},
		{
			Name:     "Punctuation groups",
			Input:    "Punctuation tests are ... kinda boring ,what do you think ?",
			Expected: "Punctuation tests are... kinda boring, what do you think?",
		},
		{
			Name:     "Complex example from requirements",
			Input:    "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			Expected: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
	}
}