package troop_test

import (
	"fmt"
	"testing"

	"github.com/brainboxweb/advent-2022/day11/troop"
	"github.com/stretchr/testify/assert"
)

func TestRounds(t *testing.T) {
	worryDivisor := 3
	tests := []struct {
		roundsCount int
		expected0   []int
		expected1   []int
	}{
		{
			1,
			[]int{20, 23, 27, 26},
			[]int{2080, 25, 167, 207, 401, 1046},
		},
		{
			2,
			[]int{695, 10, 71, 135, 350},
			[]int{43, 49, 58, 55, 362},
		},
		{
			3,
			[]int{16, 18, 21, 20, 122},
			[]int{1468, 22, 150, 286, 739},
		},
		{
			4,
			[]int{491, 9, 52, 97, 248, 34},
			[]int{39, 45, 43, 258},
		},
		{
			5,
			[]int{15, 17, 16, 88, 1037},
			[]int{20, 110, 205, 524, 72},
		},
		{
			6,
			[]int{8, 70, 176, 26, 34},
			[]int{481, 32, 36, 186, 2190},
		},
		{
			15,
			[]int{83, 44, 8, 184, 9, 20, 26, 102},
			[]int{110, 36},
		},
		{
			20,
			[]int{10, 12, 14, 26, 34},
			[]int{245, 93, 53, 199, 115},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Round %d", tt.roundsCount), func(t *testing.T) {
			mm := getTroop(worryDivisor)
			mm.PlayRounds(tt.roundsCount)
			assert.Equal(t, tt.expected0, mm.Monkeys[0].Items)
			assert.Equal(t, tt.expected1, mm.Monkeys[1].Items)
		})
	}
}

func TestMonkeyBusiness(t *testing.T) {
	worryDivisor := 3
	tests := []struct {
		roundsCount int
		expected    int
	}{
		{
			1,
			20, // 4 * 5
		},
		{
			2,
			100, // 10 * 10
		},
		{
			3,
			225, // 15 * 15
		},
		{
			20,
			10605, // 105 * 101
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Round %d", tt.roundsCount), func(t *testing.T) {
			mm := getTroop(worryDivisor)
			mm.PlayRounds(tt.roundsCount)
			result := mm.MonkeyBusiness()

			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMonkeyBusinessTwo(t *testing.T) {
	worryDivisor := 3
	tests := []struct {
		roundsCount int
		expected    int
	}{
		{
			20,
			88208,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Round %d", tt.roundsCount), func(t *testing.T) {
			mm := getTroopBig(worryDivisor)
			mm.PlayRounds(tt.roundsCount)
			result := mm.MonkeyBusiness()
			assert.Equal(t, tt.expected, result)
		})
	}
}

// --- Helpers

func getTroop(worryDivisor int) *troop.Troop {
	mm := troop.New(worryDivisor)
	mm.AddMonkey(0, []int{79, 98}, "*", "19", 23, 2, 3)
	mm.AddMonkey(1, []int{54, 65, 75, 74}, "+", "6", 19, 2, 0)
	mm.AddMonkey(2, []int{79, 60, 97}, "*", "old", 13, 1, 3)
	mm.AddMonkey(3, []int{74}, "+", "3", 17, 0, 1)

	return mm
}

func getTroopBig(worryDivisor int) *troop.Troop {
	mm := troop.New(worryDivisor) // rename
	mm.AddMonkey(0, []int{71, 86}, "*", "13", 19, 6, 7)
	mm.AddMonkey(1, []int{66, 50, 90, 53, 88, 85}, "+", "3", 2, 5, 4)
	mm.AddMonkey(2, []int{97, 54, 89, 62, 84, 80, 63}, "+", "6", 13, 4, 1)
	mm.AddMonkey(3, []int{82, 97, 56, 92}, "+", "2", 5, 6, 0)
	mm.AddMonkey(4, []int{50, 99, 67, 61, 86}, "*", "old", 7, 5, 3)
	mm.AddMonkey(5, []int{61, 66, 72, 55, 64, 53, 72, 63}, "+", "4", 11, 3, 0)
	mm.AddMonkey(6, []int{59, 79, 63}, "*", "7", 17, 2, 7)
	mm.AddMonkey(7, []int{55}, "+", "7", 3, 2, 1)

	return mm
}
