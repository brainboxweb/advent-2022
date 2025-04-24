package day1_test

import (
	"testing"

	"github.com/brainboxweb/advent-2022/day1"
	"github.com/brainboxweb/advent-2022/helpers"

	"github.com/stretchr/testify/assert"
)

const dataPath = "../testdata"

func TestDay1(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "/day1.txt",
			72478,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetData(tt.dataFile)
			result := day1.Part1(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDay1_2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "/day1_test.txt",
			45000,
		},
		{
			dataPath + "/day1.txt",
			210367,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetData(tt.dataFile)
			result := day1.Part2(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}
