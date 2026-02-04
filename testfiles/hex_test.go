package testfiles

import (
	"go-reloaded/internal"
	"testing"
)

func TestApplyHex(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		want   []string
		result bool
	}{
		{"valid hex", []string{"1E"}, []string{"30"}, true},
		{"valid hex lowercase", []string{"ff"}, []string{"255"}, true},
		{"valid hex uppercase", []string{"FF"}, []string{"255"}, true},
		{"invalid hex", []string{"ZZ"}, []string{"ZZ"}, false},
		{"empty slice", []string{}, []string{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make([]string, len(tt.input))
			copy(result, tt.input)
			got := internal.ApplyHex(result)
			if got != tt.result {
				t.Errorf("ApplyHex() = %v, want %v", got, tt.result)
			}
			if got && result[0] != tt.want[0] {
				t.Errorf("ApplyHex() result = %v, want %v", result[0], tt.want[0])
			}
		})
	}
}
