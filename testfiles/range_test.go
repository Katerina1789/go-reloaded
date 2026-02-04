package testfiles

import (
	"go-reloaded/internal"
	"testing"
)

func TestIsRangeRule(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"valid up range", "(up, 2)", true},
		{"valid low range", "(low, 3)", true},
		{"valid cap range", "(cap, 5)", true},
		{"no space", "(up,2)", true},
		{"invalid rule", "(up)", false},
		{"invalid format", "up, 2", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.IsRangeRule(tt.input); got != tt.want {
				t.Errorf("IsRangeRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplyRangeRule(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		rule  string
		want  []string
	}{
		{"up range", []string{"hello", "world"}, "(up, 2)", []string{"HELLO", "WORLD"}},
		{"low range", []string{"HELLO", "WORLD"}, "(low, 2)", []string{"hello", "world"}},
		{"cap range", []string{"hello", "world"}, "(cap, 2)", []string{"Hello", "World"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make([]string, len(tt.input))
			copy(result, tt.input)
			internal.ApplyRangeRule(result, tt.rule)
			for i := range result {
				if result[i] != tt.want[i] {
					t.Errorf("ApplyRangeRule() = %v, want %v", result, tt.want)
					break
				}
			}
		})
	}
}
