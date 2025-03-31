package day4_test

import (
	"testing"

	"github.com/brainboxweb/advent-2022/day4"
	"github.com/brainboxweb/advent-2022/helpers"
	"github.com/stretchr/testify/assert"
)

const dataPath = "../data/"

func TestPart1(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day4_test.txt",
			2,
		},
		{
			dataPath + "day4.txt",
			496,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day4.Part1(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day4_test.txt",
			4,
		},
		{
			dataPath + "day4.txt",
			847,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day4.Part2(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}
