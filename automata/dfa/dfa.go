package dfa

import "github.com/paulgriffiths/gods/sets"

// Dfa implements a deterministic finite automaton.
type Dfa struct {
	Q      int            // Number of states
	S      []rune         // Alphabet
	D      []map[rune]int // Transition function
	Start  int            // Start state
	Accept sets.SetInt    // Set of accepting states
}

// Accepts returns true if the DFA accepts the provided string.
func (d Dfa) Accepts(input string) bool {
	currentState := d.Start
	ok := false

	for _, letter := range input {
		currentState, ok = d.D[currentState][letter]
		if !ok {
			return false
		}
	}

	return d.Accept.Contains(currentState)
}
