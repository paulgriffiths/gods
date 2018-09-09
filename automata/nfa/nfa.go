package nfa

import "github.com/paulgriffiths/gods/sets"

// Nfa implements a nondeterministic finite automaton.
type Nfa struct {
	Q      int                    // Number of states
	S      []rune                 // Alphabet
	D      []map[rune]sets.SetInt // Transition function
	Start  int                    // Start state
	Accept sets.SetInt            // Set of accepting states
}

// Accepts returns true if the NFA accepts the provided string.
func (n Nfa) Accepts(input string) bool {
	current := n.Eclosure(sets.NewSetInt(n.Start))

	for _, letter := range input {
		next := sets.NewSetInt()
		for _, state := range current.Elements() {
			if p, ok := n.D[state][letter]; ok {
				next = next.Union(n.Eclosure(p))
			}
		}
		if current.IsEmpty() {
			return false
		}
		current = next
	}

	return !n.Accept.Intersection(current).IsEmpty()
}

// Eclosure returns the set of states reachable from the provided
// set of states on e-transitions alone, including the provided
// set of states itself.
func (n Nfa) Eclosure(s sets.SetInt) sets.SetInt {
	current := s
	ec := s
	prevLen := -1

	for ec.Length() != prevLen {
		prevLen = ec.Length()
		next := sets.NewSetInt()
		for _, state := range current.Elements() {
			if eStates, ok := n.D[state][0]; ok {
				ec = ec.Union(eStates)
				next = next.Union(eStates)
			}
		}
		current = next
	}

	return ec
}
