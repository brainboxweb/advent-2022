package day9_test

import (
	"testing"

	"github.com/brainboxweb/advent-2022/day9"
	"github.com/brainboxweb/advent-2022/helpers"
	"github.com/stretchr/testify/assert"
)

const dataPath = "../data/"

func TestDay9(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day9_test.txt",
			13,
		},
		{
			dataPath + "day9.txt",
			5695,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataset := helpers.GetDataString(tt.dataFile)
			result := day9.Day9(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}
