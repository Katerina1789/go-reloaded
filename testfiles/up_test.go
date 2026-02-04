package testfiles

import (
	"go-reloaded/internal"
	"testing"
)

func TestApplyUp(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		count int
		want  []string
	}{
		{"single word", []string{"hello"}, 1, []string{"HELLO"}},
		{"two words", []string{"hello", "world"}, 2, []string{"HELLO", "WORLD"}},
		{"count exceeds length", []string{"hello"}, 5, []string{"HELLO"}},
		{"empty slice", []string{}, 1, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make([]string, len(tt.input))
			copy(result, tt.input)
			internal.ApplyUp(result, tt.count)
			for i := range result {
				if result[i] != tt.want[i] {
					t.Errorf("ApplyUp() = %v, want %v", result, tt.want)
					break
				}
			}
		})
	}
}
