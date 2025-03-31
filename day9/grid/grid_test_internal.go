package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//revive:disable:function-length

func TestMoveRight(t *testing.T) {
	tests := []struct {
		name               string
		headStart          Point
		tailStart          Point
		moveCount          int
		expectedTailPoints []Point
	}{
		{
			name:               "Start below",
			headStart:          Point{X: 2, Y: 2},
			tailStart:          Point{X: 2, Y: 1},
			moveCount:          2,
			expectedTailPoints: []Point{{X: 3, Y: 2}},
		},
		{
			name:               "Start above",
			headStart:          Point{X: 2, Y: 2},
			tailStart:          Point{X: 2, Y: 3},
			moveCount:          2,
			expectedTailPoints: []Point{{X: 3, Y: 2}},
		},
		{
			name:               "Start below... but only one move",
			headStart:          Point{X: 2, Y: 2},
			tailStart:          Point{X: 2, Y: 1},
			moveCount:          1,
			expectedTailPoints: nil,
		},
		{
			name:               "Start above... but only one move",
			headStart:          Point{X: 2, Y: 2},
			tailStart:          Point{X: 2, Y: 3},
			moveCount:          1,
			expectedTailPoints: nil,
		},
		{
			name:               "Start to the left",
			headStart:          Point{X: 2, Y: 2}, // to 4,4
			tailStart:          Point{X: 1, Y: 2},
			moveCount:          2,
			expectedTailPoints: []Point{{X: 2, Y: 2}, {X: 3, Y: 2}},
		},
		{
			name:               "Start to the right",
			headStart:          Point{X: 2, Y: 2},
			tailStart:          Point{X: 3, Y: 2},
			moveCount:          2,
			expectedTailPoints: nil,
		},
		{
			name:               "Start above left",
			headStart:          Point{X: 2, Y: 2},
			tailStart:          Point{X: 1, Y: 3},
			moveCount:          2,
			expectedTailPoints: []Point{{X: 2, Y: 2}, {X: 3, Y: 2}},
		},
		{
			name:               "Start right",
			headStart:          Point{X: 2, Y: 2},
			tailStart:          Point{X: 3, Y: 2},
			moveCount:          2,
			expectedTailPoints: nil,
		},
		{
			name:               "Start co-located",
			headStart:          Point{X: 2, Y: 2},
			tailStart:          Point{X: 2, Y: 2},
			moveCount:          2,
			expectedTailPoints: []Point{{X: 3, Y: 2}},
		},
	}

	for _, tt := range tests {
		tg := New()
		tg.Head = tt.headStart
		tg.Tail = tt.tailStart
		t.Run(tt.name, func(t *testing.T) {
			result := tg.moveRight(tt.moveCount)
			assert.Equal(t, tt.expectedTailPoints, result)
		})
	}
}

func TestMoveLeft(t *testing.T) {
	tests := []struct {
		name               string
		headStart          Point
		tailStart          Point
		moveCount          int
		expectedTailPoints []Point
	}{
		{
			name:               "Start below",
			headStart:          Point{X: 4, Y: 4},
			tailStart:          Point{X: 4, Y: 3},
			moveCount:          2,
			expectedTailPoints: []Point{{X: 3, Y: 4}},
		},
		{
			name:               "Start above",
			headStart:          Point{X: 4, Y: 4},
			tailStart:          Point{X: 4, Y: 5},
			moveCount:          2,
			expectedTailPoints: []Point{{X: 3, Y: 4}},
		},
		{
			name:               "Start below... but only one move",
			headStart:          Point{X: 4, Y: 4},
			tailStart:          Point{X: 4, Y: 3},
			moveCount:          1,
			expectedTailPoints: nil,
		},
		{
			name:               "Start above... but only one move",
			headStart:          Point{X: 4, Y: 4},
			tailStart:          Point{X: 4, Y: 5},
			moveCount:          1,
			expectedTailPoints: nil,
		},
		{
			name:               "Start to the left",
			headStart:          Point{X: 4, Y: 4},
			tailStart:          Point{X: 3, Y: 4},
			moveCount:          2,
			expectedTailPoints: nil,
		},
		{
			name:               "Start to the right",
			headStart:          Point{X: 4, Y: 4}, // ends 2, 4
			tailStart:          Point{X: 5, Y: 4},
			moveCount:          2,
			expectedTailPoints: []Point{{X: 4, Y: 4}, {X: 3, Y: 4}},
		},
		{
			name:               "Start above left",
			headStart:          Point{X: 4, Y: 4}, // 2,2
			tailStart:          Point{X: 3, Y: 5},
			moveCount:          2,
			expectedTailPoints: nil,
		},
		{
			name:               "Start right",
			headStart:          Point{X: 4, Y: 4},
			tailStart:          Point{X: 5, Y: 4},
			moveCount:          2,
			expectedTailPoints: []Point{{X: 4, Y: 4}, {X: 3, Y: 4}},
		},
		{
			name:               "Start co-located",
			headStart:          Point{X: 4, Y: 4},
			tailStart:          Point{X: 4, Y: 4},
			moveCount:          2,
			expectedTailPoints: []Point{{X: 3, Y: 4}},
		},
	}

	for _, tt := range tests {
		tg := New()
		tg.Head = tt.headStart
		tg.Tail = tt.tailStart
		t.Run(tt.name, func(t *testing.T) {
			result := tg.moveLeft(tt.moveCount)
			assert.Equal(t, tt.expectedTailPoints, result)
		})
	}
}

func TestMoveUp(t *testing.T) {
	tests := []struct {
		name               string
		headStart          Point
		tailStart          Point
		moveCount          int
		expectedTailPoints []Point
	}{
		{
			name:               "Start below",
			headStart:          Point{X: 2, Y: 2}, // 2, 4
			tailStart:          Point{X: 2, Y: 1},
			moveCount:          2,
			expectedTailPoints: []Point{{X: 2, Y: 2}, {X: 2, Y: 3}},
		},
		{
			name:               "Start above",
			headStart:          Point{X: 2, Y: 2},
			tailStart:          Point{X: 2, Y: 3},
			moveCount:          2,
			expectedTailPoints: nil,
		},
		{
			name:               "Start below... but only one move",
			headStart:          Point{X: 2, Y: 2},
			tailStart:          Point{X: 2, Y: 1},
			moveCount:          1,
			expectedTailPoints: []Point{{X: 2, Y: 2}},
		},
		{
			name:               "Start above... but only one move",
			headStart:          Point{X: 2, Y: 2},
			tailStart:          Point{X: 2, Y: 3},
			moveCount:          1,
			expectedTailPoints: nil,
		},
		{
			name:               "Start co-located",
			headStart:          Point{X: 2, Y: 2},
			tailStart:          Point{X: 2, Y: 2},
			moveCount:          2,
			expectedTailPoints: []Point{{X: 2, Y: 3}},
		},
	}

	for _, tt := range tests {
		tg := New()
		tg.Head = tt.headStart
		tg.Tail = tt.tailStart
		t.Run(tt.name, func(t *testing.T) {
			result := tg.moveUp(tt.moveCount)
			assert.Equal(t, tt.expectedTailPoints, result)
		})
	}
}

func TestMoveDown(t *testing.T) {
	tests := []struct {
		name               string
		headStart          Point
		tailStart          Point
		moveCount          int
		expectedTailPoints []Point
	}{
		{
			name:               "Start below",
			headStart:          Point{X: 4, Y: 4},
			tailStart:          Point{X: 4, Y: 3},
			moveCount:          2,
			expectedTailPoints: nil,
		},
		{
			name:               "Start above",
			headStart:          Point{X: 4, Y: 4},
			tailStart:          Point{X: 4, Y: 5},
			moveCount:          2,
			expectedTailPoints: []Point{{X: 4, Y: 4}, {X: 4, Y: 3}},
		},
	}

	for _, tt := range tests {
		tg := New()
		tg.Head = tt.headStart
		tg.Tail = tt.tailStart
		t.Run(tt.name, func(t *testing.T) {
			result := tg.moveDown(tt.moveCount)
			assert.Equal(t, tt.expectedTailPoints, result)
		})
	}
}
