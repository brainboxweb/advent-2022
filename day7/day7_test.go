package day7_test

import (
	"testing"

	"github.com/brainboxweb/advent-2022/day7"
	"github.com/brainboxweb/advent-2022/helpers"
	"github.com/stretchr/testify/assert"
)

const dataPath = "../testdata"

func TestPart1(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "/day7_test.txt",
			95437,
		},
		{
			dataPath + "/day7.txt",
			1501149,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day7.Part1(dataSet)
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
			dataPath + "/day7_test.txt",
			24933642,
		},
		{
			dataPath + "/day7.txt",
			10096985,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day7.Part2(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}
