package testfiles

import (
	"go-reloaded/internal"
	"testing"
)

func TestApplyLow(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		count int
		want  []string
	}{
		{"single word", []string{"HELLO"}, 1, []string{"hello"}},
		{"two words", []string{"HELLO", "WORLD"}, 2, []string{"hello", "world"}},
		{"count exceeds length", []string{"HELLO"}, 5, []string{"hello"}},
		{"empty slice", []string{}, 1, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make([]string, len(tt.input))
			copy(result, tt.input)
			internal.ApplyLow(result, tt.count)
			for i := range result {
				if result[i] != tt.want[i] {
					t.Errorf("ApplyLow() = %v, want %v", result, tt.want)
					break
				}
			}
		})
	}
}
