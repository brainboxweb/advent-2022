package day5_test

import (
	"testing"

	"github.com/brainboxweb/advent-2022/day5"
	"github.com/brainboxweb/advent-2022/helpers"
	"github.com/stretchr/testify/assert"
)

const dataPath = "../data/"

func TestPart1(t *testing.T) {
	tests := []struct {
		dataFile string
		expected string
	}{
		{
			dataPath + "day5_test.txt",
			"CMZ",
		},
		{
			dataPath + "day5.txt",
			"QGTHFZBHV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day5.Part1(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected string
	}{
		{
			dataPath + "day5_test.txt",
			"MCD",
		},
		{
			dataPath + "day5.txt",
			"MGDMPSZTM",
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day5.Part2(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}
