package internal

import "strings" // For joining result words

// State represents the current FSM state
type State int

// Define the two states our FSM can be in
const (
	Normal    State = iota // Default state - processing normal text
	QuoteOpen              // Inside quotes - same rules but tracks quote context
)

// FSM holds the current state and accumulated result
type FSM struct {
	State  State    // Current state (Normal or QuoteOpen)
	Result []string // Words processed so far
}

// NewFSM creates a new FSM starting in Normal state
func NewFSM() *FSM {
	return &FSM{State: Normal, Result: make([]string, 0)}
}

// ProcessWord handles each word based on current FSM state
func (fsm *FSM) ProcessWord(word string) {
	switch fsm.State {
	case Normal:
		fsm.handleNormal(word) // Process word in normal context
	case QuoteOpen:
		fsm.handleQuote(word) // Process word inside quotes
	}
}

// handleNormal processes words in normal state
func (fsm *FSM) handleNormal(word string) {
	// Try hex conversion - only consume token if conversion succeeds
	if word == "(hex)" && len(fsm.Result) > 0 {
		if ApplyHex(fsm.Result) { // Returns true if conversion worked
			return // Don't add "(hex)" to result
		}
	}
	// Try binary conversion - only consume token if conversion succeeds
	if word == "(bin)" && len(fsm.Result) > 0 {
		if ApplyBin(fsm.Result) { // Returns true if conversion worked
			return // Don't add "(bin)" to result
		}
	}
	// Apply uppercase to previous word
	if word == "(up)" && len(fsm.Result) > 0 {
		ApplyUp(fsm.Result, 1) // Transform 1 word
		return                 // Don't add "(up)" to result
	}
	// Apply lowercase to previous word
	if word == "(low)" && len(fsm.Result) > 0 {
		ApplyLow(fsm.Result, 1) // Transform 1 word
		return                  // Don't add "(low)" to result
	}
	// Apply capitalization to previous word
	if word == "(cap)" && len(fsm.Result) > 0 {
		ApplyCap(fsm.Result, 1) // Transform 1 word
		return                  // Don't add "(cap)" to result
	}
	// Check for range rules like "(up, 2)"
	if IsRangeRule(word) {
		ApplyRangeRule(fsm.Result, word) // Transform multiple words
		return                           // Don't add range rule to result
	}
	// Handle quote opening - switch to QuoteOpen state
	if word == "'" {
		fsm.State = QuoteOpen // Change state to track quote context
	}
	// Add word to result (includes quotes and regular words)
	fsm.Result = append(fsm.Result, word)
}

// handleQuote processes words inside quotes (same rules as normal)
func (fsm *FSM) handleQuote(word string) {
	// Same transformation rules apply inside quotes
	if word == "(hex)" && len(fsm.Result) > 0 {
		if ApplyHex(fsm.Result) { // Only consume if conversion works
			return
		}
	}
	if word == "(bin)" && len(fsm.Result) > 0 {
		if ApplyBin(fsm.Result) { // Only consume if conversion works
			return
		}
	}
	if word == "(up)" && len(fsm.Result) > 0 {
		ApplyUp(fsm.Result, 1) // Apply uppercase inside quotes
		return
	}
	if word == "(low)" && len(fsm.Result) > 0 {
		ApplyLow(fsm.Result, 1) // Apply lowercase inside quotes
		return
	}
	if word == "(cap)" && len(fsm.Result) > 0 {
		ApplyCap(fsm.Result, 1) // Apply capitalization inside quotes
		return
	}
	if IsRangeRule(word) {
		ApplyRangeRule(fsm.Result, word) // Apply range rules inside quotes
		return
	}
	// Handle quote closing - return to Normal state
	if word == "'" {
		fsm.State = Normal // Switch back to normal state
	}
	// Add word to result
	fsm.Result = append(fsm.Result, word)
}

// GetOutput returns final processed text with post-processing
func (fsm *FSM) GetOutput() string {
	text := strings.Join(fsm.Result, " ") // Join all words with spaces
	text = FixPunctuation(text)           // Fix punctuation spacing
	text = FixQuotes(text)                // Fix quote formatting
	text = FixArticles(text)              // Fix a/an articles
	return text                           // Return final result
}