package testfiles

import (
	"go-reloaded/internal"
	"testing"
)

func TestApplyBin(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		want   []string
		result bool
	}{
		{"valid binary", []string{"101"}, []string{"5"}, true},
		{"valid binary 10", []string{"10"}, []string{"2"}, true},
		{"valid binary 1111", []string{"1111"}, []string{"15"}, true},
		{"invalid binary", []string{"22"}, []string{"22"}, false},
		{"empty slice", []string{}, []string{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make([]string, len(tt.input))
			copy(result, tt.input)
			got := internal.ApplyBin(result)
			if got != tt.result {
				t.Errorf("ApplyBin() = %v, want %v", got, tt.result)
			}
			if got && result[0] != tt.want[0] {
				t.Errorf("ApplyBin() result = %v, want %v", result[0], tt.want[0])
			}
		})
	}
}
