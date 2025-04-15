package day10_test

import (
	"testing"

	"github.com/brainboxweb/advent-2022/day10"
	"github.com/brainboxweb/advent-2022/helpers"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		data     []string
		dataFile string
		expected int
	}{
		{
			nil,
			"../data/day10_test.txt",
			13140,
		},
		{
			[]string{},
			"../data/day10.txt",
			15120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			var dataset []string
			if len(tt.data) == 0 {
				dataset = helpers.GetDataString(tt.dataFile)
			} else {
				dataset = tt.data
			}
			result := day10.Part1(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		data     []string
		dataFile string
		expected string
	}{
		{
			nil,
			"../data/day10_test.txt",
			expectedOne,
		},
		{
			[]string{},
			"../data/day10.txt",
			expectedTwo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			var dataset []string
			if len(tt.data) == 0 {
				dataset = helpers.GetDataString(tt.dataFile)
			} else {
				dataset = tt.data
			}
			result := day10.Part2(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}

const expectedOne = `##..##..##..##..##..##..##..##..##..##.
###...###...###...###...###...###...###
####....####....####....####....####...
#####.....#####.....#####.....#####....
######......######......######......###
#######.......#######.......#######....`

const expectedTwo = `###..#..#.###....##.###..###..#.....##.
#..#.#.#..#..#....#.#..#.#..#.#....#..#
#..#.##...#..#....#.###..#..#.#....#..#
###..#.#..###.....#.#..#.###..#....####
#.#..#.#..#....#..#.#..#.#....#....#..#
#..#.#..#.#.....##..###..#....####.#..#`

// RKPJBPLA
