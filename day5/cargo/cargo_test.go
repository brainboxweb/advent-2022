package cargo_test

import (
	"testing"

	"github.com/brainboxweb/advent-2022/day5/cargo"
	"github.com/stretchr/testify/assert"
)

func TestAddCrates(t *testing.T) {
	tests := []struct {
		data     []string
		expected []string
	}{
		{
			data:     []string{"Z", "N", "D"},
			expected: []string{"Z", "N", "D"},
		},
		{
			data:     []string{"M", "C"},
			expected: []string{"M", "C"},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			stacks := &cargo.Stacks{}
			stacks.AddStack()
			for _, crateName := range tt.data {
				stacks.AddCrateToStack(0, crateName)
			}
			assert.Equal(t, tt.expected, stacks.GetCrates(0))
		})
	}
}

func TestMoves(t *testing.T) {
	stacks := &cargo.Stacks{}
	stacks.AddStack()
	stacks.AddCrateToStack(0, "Z")
	stacks.AddCrateToStack(0, "N")

	stacks.AddStack()
	stacks.AddCrateToStack(1, "M")
	stacks.AddCrateToStack(1, "C")
	stacks.AddCrateToStack(1, "D")

	stacks.AddStack()
	stacks.AddCrateToStack(2, "P")

	// move 1 from 2 to 1
	// move 3 from 1 to 3
	// move 2 from 2 to 1
	// move 1 from 1 to 2

	stacks.ApplyMove(cargo.Move{1, 1, 0})
	stacks.ApplyMove(cargo.Move{3, 0, 2})
	stacks.ApplyMove(cargo.Move{2, 1, 0})
	stacks.ApplyMove(cargo.Move{1, 0, 1})

	assert.Equal(t, []string{"C"}, stacks.GetCrates(0))
	assert.Equal(t, []string{"M"}, stacks.GetCrates(1))
	assert.Equal(t, []string{"P", "D", "N", "Z"}, stacks.GetCrates(2))

	assert.Equal(t, "CMZ", stacks.GetTop())
}

func TestMovesReverse(t *testing.T) {
	stacks := &cargo.Stacks{}
	stacks.AddStack()
	stacks.AddCrateToStack(0, "Z")
	stacks.AddCrateToStack(0, "N")

	stacks.AddStack()
	stacks.AddCrateToStack(1, "M")
	stacks.AddCrateToStack(1, "C")
	stacks.AddCrateToStack(1, "D")

	stacks.AddStack()
	stacks.AddCrateToStack(2, "P")

	// move 1 from 2 to 1
	// move 3 from 1 to 3
	// move 2 from 2 to 1
	// move 1 from 1 to 2

	stacks.ApplyMoveWithReverse(cargo.Move{1, 1, 0})
	stacks.ApplyMoveWithReverse(cargo.Move{3, 0, 2})
	stacks.ApplyMoveWithReverse(cargo.Move{2, 1, 0})
	stacks.ApplyMoveWithReverse(cargo.Move{1, 0, 1})

	assert.Equal(t, []string{"M"}, stacks.GetCrates(0))
	assert.Equal(t, []string{"C"}, stacks.GetCrates(1))
	assert.Equal(t, []string{"P", "Z", "N", "D"}, stacks.GetCrates(2))

	assert.Equal(t, "MCD", stacks.GetTop())
}
