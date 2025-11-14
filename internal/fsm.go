package internal

import "strings"

type State int

const (
	Normal State = iota
	QuoteOpen
)

type FSM struct {
	State  State
	Result []string
}

func NewFSM() *FSM {
	return &FSM{State: Normal, Result: make([]string, 0)}
}

func (fsm *FSM) ProcessWord(word string) {
	switch fsm.State {
	case Normal:
		fsm.handleNormal(word)
	case QuoteOpen:
		fsm.handleQuote(word)
	}
}

func (fsm *FSM) handleNormal(word string) {
	if word == "(hex)" && len(fsm.Result) > 0 {
		ApplyHex(fsm.Result)
		return
	}
	if word == "(bin)" && len(fsm.Result) > 0 {
		ApplyBin(fsm.Result)
		return
	}
	if word == "(up)" && len(fsm.Result) > 0 {
		ApplyUp(fsm.Result, 1)
		return
	}
	if word == "(low)" && len(fsm.Result) > 0 {
		ApplyLow(fsm.Result, 1)
		return
	}
	if word == "(cap)" && len(fsm.Result) > 0 {
		ApplyCap(fsm.Result, 1)
		return
	}
	if IsRangeRule(word) {
		ApplyRangeRule(fsm.Result, word)
		return
	}
	if word == "'" {
		fsm.State = QuoteOpen
	}
	fsm.Result = append(fsm.Result, word)
}

func (fsm *FSM) handleQuote(word string) {
	if word == "(hex)" && len(fsm.Result) > 0 {
		ApplyHex(fsm.Result)
		return
	}
	if word == "(bin)" && len(fsm.Result) > 0 {
		ApplyBin(fsm.Result)
		return
	}
	if word == "(up)" && len(fsm.Result) > 0 {
		ApplyUp(fsm.Result, 1)
		return
	}
	if word == "(low)" && len(fsm.Result) > 0 {
		ApplyLow(fsm.Result, 1)
		return
	}
	if word == "(cap)" && len(fsm.Result) > 0 {
		ApplyCap(fsm.Result, 1)
		return
	}
	if IsRangeRule(word) {
		ApplyRangeRule(fsm.Result, word)
		return
	}
	if word == "'" {
		fsm.State = Normal
	}
	fsm.Result = append(fsm.Result, word)
}

func (fsm *FSM) GetOutput() string {
	text := strings.Join(fsm.Result, " ")
	text = FixPunctuation(text)
	text = FixQuotes(text)
	text = FixArticles(text)
	return text
}