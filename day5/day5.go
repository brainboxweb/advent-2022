package day5

import (
	"strconv"
	"strings"

	"github.com/brainboxweb/advent-2022/day5/cargo"
	"github.com/brainboxweb/advent-2022/helpers"
)

func Part1(data []string) string {
	crateStrings, moveStrings := parse(data)
	stacks := buildStacks(crateStrings)

	// Moves
	moves := prepMoves(moveStrings)
	for _, move := range moves {
		stacks.ApplyMove(move)
	}
	top := stacks.GetTop()

	return top
}

func Part2(data []string) string {
	crateStrings, moveStrings := parse(data)
	stacks := buildStacks(crateStrings)

	// Moves
	moves := prepMoves(moveStrings)
	for _, move := range moves {
		stacks.ApplyMoveWithReverse(move) // Rename
	}
	top := stacks.GetTop()

	return top
}

func buildStacks(crateStrings []string) cargo.Stacks {
	crateData := prepStacks(crateStrings)
	stacks := cargo.Stacks{}
	for range crateData[0] {
		stacks.AddStack()
	}
	for _, items := range crateData {
		for i, crate := range items {
			if crate != " " {
				stacks.AddCrateToStack(i, crate)
			}
		}
	}
	return stacks
}

func prepStacks(crateStrings []string) [][]string {
	crateData := make([][]string, 0)
	for _, line := range crateStrings {
		parts := []string{}
		max := len(line) / 4
		if len(line)%4 != 0 {
			max++
		}
		for i := 0; i < max; i++ {
			parts = append(parts, string(line[i*4+1]))
		}
		crateData = append(crateData, parts)
	}
	// drop the last element
	crateData = crateData[:len(crateData)-1]
	crateData = helpers.ReverseSliceOfSlices(crateData)

	return crateData
}

func prepMoves(moveStrings []string) []cargo.Move {
	moves := []cargo.Move{}
	for _, line := range moveStrings {
		parts := strings.Split(line, " ")
		qty, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])
		move := cargo.Move{
			Quant: qty,
			From:  from - 1, // to zero-based
			To:    to - 1,   // to zero-based
		}
		moves = append(moves, move)
	}

	return moves
}

func parse(input []string) (crateStrings, moveStrings []string) {
	isCrate := true
	for _, line := range input {
		if line == "" {
			isCrate = false
			continue
		}
		if isCrate {
			crateStrings = append(crateStrings, line)
			continue
		}
		moveStrings = append(moveStrings, line)
	}

	return crateStrings, moveStrings
}
