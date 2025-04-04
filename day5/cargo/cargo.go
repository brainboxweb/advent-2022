package cargo

import (
	"github.com/brainboxweb/advent-2022/helpers"
)

type Move struct {
	Quant int
	From  int
	To    int
}

type Stacks struct {
	Stacks []Stack
}

func (ss *Stacks) AddStack() {
	ss.Stacks = append(ss.Stacks, Stack{})
}

func (ss *Stacks) AddCrateToStack(stackID int, crate string) {
	ss.Stacks[stackID].addCrate(crate)
}

// For testing
func (ss *Stacks) GetCrates(stackID int) []string {
	return ss.Stacks[stackID].Crates
}

func (ss *Stacks) ApplyMove(move Move) {
	for i := 0; i < move.Quant; i++ {
		inTransit := ss.Stacks[move.From].removeCrate()
		ss.Stacks[move.To].addCrate(inTransit)
	}
}

func (ss *Stacks) ApplyMoveWithReverse(move Move) {
	inTransits := []string{}
	for i := 0; i < move.Quant; i++ {
		inTransit := ss.Stacks[move.From].removeCrate()
		inTransits = append(inTransits, inTransit)
	}
	// Put them back in reverse order
	inTransits = helpers.ReverseSlice(inTransits)
	for _, crate := range inTransits {
		ss.Stacks[move.To].addCrate(crate)
	}
}

func (ss *Stacks) GetTop() string {
	ret := ""
	for _, stack := range ss.Stacks {
		top := stack.getTop()
		ret += top
	}

	return ret
}

type Stack struct {
	Crates []string
}

func (s *Stack) addCrate(crate string) {
	s.Crates = append(s.Crates, crate)
}

func (s *Stack) removeCrate() string {
	inTransit := s.Crates[len(s.Crates)-1]
	s.Crates = s.Crates[:len(s.Crates)-1]

	return inTransit
}

func (s *Stack) getTop() string {
	return s.Crates[len(s.Crates)-1]
}
