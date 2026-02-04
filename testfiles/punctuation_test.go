package testfiles

import (
	"go-reloaded/internal"
	"testing"
)

func TestFixPunctuation(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"space before comma", "hello , world", "hello, world"},
		{"space before period", "hello . world", "hello. world"},
		{"grouped dots", "hello ...", "hello..."},
		{"exclamation", "hello !", "hello!"},
		{"question", "hello ?", "hello?"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.FixPunctuation(tt.input); got != tt.want {
				t.Errorf("FixPunctuation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFixQuotes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"single word", "' hello '", "'hello'"},
		{"multiple words", "' hello world '", "'hello world'"},
		{"no spaces", "'hello'", "'hello'"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.FixQuotes(tt.input); got != tt.want {
				t.Errorf("FixQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFixArticles(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"a to an vowel", "a amazing", "an amazing"},
		{"a to an h", "a house", "an house"},
		{"A to An vowel", "A amazing", "An amazing"},
		{"keep a consonant", "a dog", "a dog"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.FixArticles(tt.input); got != tt.want {
				t.Errorf("FixArticles() = %v, want %v", got, tt.want)
			}
		})
	}
}
