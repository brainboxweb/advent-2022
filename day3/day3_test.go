package day3_test

import (
	"testing"

	"github.com/brainboxweb/advent-2022/day3"
	"github.com/brainboxweb/advent-2022/helpers"
	"github.com/stretchr/testify/assert"
)

const dataPath = "../testdata"

func TestDay3(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "/day3_test.txt",
			157,
		},
		{
			dataPath + "/day3.txt",
			7848,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day3.Part1(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDay3_2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "/day3_test.txt",
			70,
		},
		{
			dataPath + "/day3.txt",
			2616,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day3.Part2(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}
