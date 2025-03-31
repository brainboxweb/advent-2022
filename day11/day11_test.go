package day11_test

import (
	"testing"

	"github.com/brainboxweb/advent-2022/day11"
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
			"../data/day11_test.txt",
			10605,
		},
		{
			[]string{},
			"../data/day11.txt",
			88208,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataset := helpers.GetDataString(tt.dataFile)
			result := day11.Part1(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// func TestPart2(t *testing.T) {
// 	tests := []struct {
// 		data     []string
// 		dataFile string
// 		expected int
// 	}{
// 		{
// 			nil,
// 			"../data/day11_test.txt",
// 			2713310158,
// 		},
// 		// {
// 		// 	[]string{},
// 		// 	"../data/day11.txt",
// 		// 	88208,
// 		// },
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.dataFile, func(t *testing.T) {
// 			var dataset []string
// 			if len(tt.data) == 0 {
// 				dataset = helpers.GetDataString(tt.dataFile)
// 			} else {
// 				dataset = tt.data
// 			}
// 			result := day11.Part2(dataset)
// 			assert.Equal(t, tt.expected, result)
// 		})
// 	}
// }
