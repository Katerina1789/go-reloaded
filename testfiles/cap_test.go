package testfiles

import (
	"go-reloaded/internal"
	"testing"
)

func TestApplyCap(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		count int
		want  []string
	}{
		{"single word", []string{"hello"}, 1, []string{"Hello"}},
		{"two words", []string{"hello", "world"}, 2, []string{"Hello", "World"}},
		{"count exceeds length", []string{"hello"}, 5, []string{"Hello"}},
		{"empty slice", []string{}, 1, []string{}},
		{"already capitalized", []string{"Hello"}, 1, []string{"Hello"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make([]string, len(tt.input))
			copy(result, tt.input)
			internal.ApplyCap(result, tt.count)
			for i := range result {
				if result[i] != tt.want[i] {
					t.Errorf("ApplyCap() = %v, want %v", result, tt.want)
					break
				}
			}
		})
	}
}
