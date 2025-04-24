package day2_test

import (
	"testing"

	"github.com/brainboxweb/advent-2022/day2"
	"github.com/brainboxweb/advent-2022/helpers"
	"github.com/stretchr/testify/assert"
)

const dataPath = "../testdata"

func TestDay2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "/day2_test.txt",
			15,
		},
		{
			dataPath + "/day2.txt",
			12458,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day2.Part1(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDay2_2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "/day2_test.txt",
			12,
		},
		{
			dataPath + "/day2.txt",
			12683,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day2.Part2(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}
