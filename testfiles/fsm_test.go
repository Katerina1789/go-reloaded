package testfiles

import (
	"go-reloaded/internal"
	"testing"
)

func TestNewFSM(t *testing.T) {
	fsm := internal.NewFSM()
	if fsm.State != internal.Normal {
		t.Errorf("NewFSM() state = %v, want %v", fsm.State, internal.Normal)
	}
	if len(fsm.Result) != 0 {
		t.Errorf("NewFSM() result length = %v, want 0", len(fsm.Result))
	}
}

func TestFSMProcessWord(t *testing.T) {
	tests := []struct {
		name   string
		words  []string
		want   string
	}{
		{"simple text", []string{"hello", "world"}, "hello world"},
		{"uppercase", []string{"hello", "(up)"}, "HELLO"},
		{"lowercase", []string{"HELLO", "(low)"}, "hello"},
		{"capitalize", []string{"hello", "(cap)"}, "Hello"},
		{"hex conversion", []string{"1E", "(hex)"}, "30"},
		{"binary conversion", []string{"101", "(bin)"}, "5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsm := internal.NewFSM()
			for _, word := range tt.words {
				fsm.ProcessWord(word)
			}
			if got := fsm.GetOutput(); got != tt.want {
				t.Errorf("FSM.GetOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSMQuoteState(t *testing.T) {
	fsm := internal.NewFSM()
	fsm.ProcessWord("'")
	if fsm.State != internal.QuoteOpen {
		t.Errorf("State after ' = %v, want %v", fsm.State, internal.QuoteOpen)
	}
	fsm.ProcessWord("hello")
	fsm.ProcessWord("'")
	if fsm.State != internal.Normal {
		t.Errorf("State after closing ' = %v, want %v", fsm.State, internal.Normal)
	}
}
