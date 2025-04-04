package day12

import (
	"sort"

	"github.com/brainboxweb/advent-2022/day12/hill"
)

const (
	startRune = 83
	endRune   = 69
	aRune     = 97
)

func Part1(data []string) int {
	sos := PrepData(data)

	visited := make(map[hill.Point]bool)
	elev := hill.Elev{SOS: sos, Visited: visited}

	maxX := len(sos)
	maxY := len(sos[0])

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			switch sos[x][y] {
			case startRune:
				elev.Start = hill.Point{X: x, Y: y}
			case endRune:
				elev.End = hill.Point{X: x, Y: y}
			}
		}
	}
	start := elev.Start

	return hill.GetStepCount(elev, start)
}

//revive:disable:cognitive-complexity

func Part2(data []string) int {
	sos := PrepData(data)
	visited := make(map[hill.Point]bool)
	elev := hill.Elev{SOS: sos, Visited: visited}
	maxX := len(sos)
	maxY := len(sos[0])

	startingPoints := []hill.Point{}

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			switch sos[x][y] {
			case startRune: // "S" - start
				startingPoints = append(startingPoints, hill.Point{X: x, Y: y})
			case aRune: // "a"
				startingPoints = append(startingPoints, hill.Point{X: x, Y: y})
			case endRune: // "E" - end
				elev.End = hill.Point{X: x, Y: y}
			}
		}
	}
	results := []int{}
	for _, start := range startingPoints {
		elev.Visited = make(map[hill.Point]bool) // NB
		res := hill.GetStepCount(elev, start)
		if res == 0 {
			continue
		}
		results = append(results, res)
	}
	sort.Ints(results)

	return results[0]
}

//revive:enable:cognitive-complexity

func PrepData(data []string) [][]rune {
	limit := len(data)
	sliceOfSlices := make([][]rune, limit) // refactor to function
	for i := 0; i < limit; i++ {
		sliceOfSlices[i] = make([]rune, limit)
		line := data[i]
		lineItems := []rune{}
		for _, theChar := range line {
			lineItems = append(lineItems, theChar)
		}
		sliceOfSlices[i] = lineItems
	}

	return transpose(sliceOfSlices)
}

// https://gist.github.com/tanaikech/5cb41424ff8be0fdf19e78d375b6adb8
func transpose(slice [][]rune) [][]rune {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]rune, xl)
	for i := range result {
		result[i] = make([]rune, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
