package dfa

// Dfa implements a deterministic finite automaton.
type Dfa struct {
	Q      int            // Number of states
	S      []rune         // Alphabet
	D      []map[rune]int // Transition function
	Start  int            // Start state
	Accept []int          // Set of accepting states
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

	for _, acceptState := range d.Accept {
		if currentState == acceptState {
			return true
		}
	}
	return false
}
