package processor

import "testing"

func TestProcessor(t *testing.T) {
	p := New()
	
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "hex conversion",
			input:    "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			expected: "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			name:     "article correction",
			input:    "There is no greater agony than bearing a untold story inside you.",
			expected: "There is no greater agony than bearing an untold story inside you.",
		},
		{
			name:     "punctuation spacing",
			input:    "Punctuation tests are ... kinda boring ,what do you think ?",
			expected: "Punctuation tests are... kinda boring, what do you think?",
		},
		{
			name:     "range transformations",
			input:    "This is so exciting (up, 2)",
			expected: "This is SO EXCITING",
		},
		{
			name:     "quotes formatting",
			input:    "I am exactly how they describe me: ' awesome '",
			expected: "I am exactly how they describe me: 'awesome'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.Process(tt.input)
			if result != tt.expected {
				t.Errorf("Process() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestStartsWithVowelOrH(t *testing.T) {
	p := New()
	
	tests := []struct {
		word     string
		expected bool
	}{
		{"amazing", true},
		{"house", true},
		{"elephant", true},
		{"umbrella", true},
		{"book", false},
		{"cat", false},
		{"dog", false},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			result := p.startsWithVowelOrH(tt.word)
			if result != tt.expected {
				t.Errorf("startsWithVowelOrH(%q) = %v, want %v", tt.word, result, tt.expected)
			}
		})
	}
}