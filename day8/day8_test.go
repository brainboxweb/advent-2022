package day8_test

import (
	"testing"

	"github.com/brainboxweb/advent-2022/day8"
	"github.com/brainboxweb/advent-2022/helpers"
	"github.com/stretchr/testify/assert"
)

const dataPath = "../testdata"

func TestDay8(t *testing.T) {
	tests := []struct {
		data     []string
		dataFile string
		expected int
	}{
		{
			[]string{"111", "101", "111"},
			"",
			8,
		},
		{
			[]string{"111", "171", "111"},
			"",
			9,
		},
		{
			[]string{"111", "111", "222"},
			"",
			8,
		},
		{
			nil,
			dataPath + "/day8_test.txt",
			21,
		},
		{
			nil,
			dataPath + "/day8.txt",
			1782,
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
			result := day8.Part1(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDay8_2(t *testing.T) {
	tests := []struct {
		data     []string
		dataFile string
		expected int
	}{
		{
			nil,
			dataPath + "/day8_test.txt",
			8,
		},
		{
			nil,
			dataPath + "/day8.txt",
			474606,
		},
	}
	for _, tt := range tests {
		dataset := []string{}
		t.Run(tt.dataFile, func(t *testing.T) {
			if len(tt.data) == 0 {
				dataset = helpers.GetDataString(tt.dataFile)
			} else {
				dataset = tt.data
			}
			result := day8.Part2(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}
