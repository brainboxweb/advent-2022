package day9

import (
	"strconv"
	"strings"

	"github.com/brainboxweb/advent-2022/day9/grid"
)

func Day9(data []string) int {
	t := grid.New()
	for _, line := range data {
		parts := strings.Split(line, " ")
		dirn := parts[0]
		moves, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("not expected")
		}
		t.Move(dirn, moves)
		// t.Plot() // for debug
	}

	return len(t.Visited)
}
