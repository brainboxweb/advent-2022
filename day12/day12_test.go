package day12_test

import (
	"testing"

	"github.com/brainboxweb/advent-2022/day12"
	"github.com/brainboxweb/advent-2022/helpers"
	"github.com/stretchr/testify/assert"
)

func TestDay12(t *testing.T) {
	tests := []struct {
		data     []string
		dataFile string
		expected int
	}{
		{
			[]string{
				"Sabqponm",
				"abcryxxl",
				"accszExk",
				"acctuvwj",
				"abdefghi",
			},
			"",
			31,
		},
		{
			[]string{},
			"../data/day12.txt",
			520,
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
			result := day12.Part1(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDay12_2(t *testing.T) {
	tests := []struct {
		data     []string
		dataFile string
		expected int
	}{
		{
			[]string{
				"Sabqponm",
				"abcryxxl",
				"accszExk",
				"acctuvwj",
				"abdefghi",
			},
			"",
			29,
		},
		{
			[]string{},
			"../data/day12.txt",
			508,
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
			result := day12.Part2(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}
