package hill

import (
	"errors"
	"fmt"
)

type Elev struct {
	SOS     [][]rune
	Start   Point
	End     Point
	Visited map[Point]bool
}

type Point struct {
	X int
	Y int
}

func (e *Elev) AlreadyVisited(p Point) bool {
	_, ok := e.Visited[p]
	return ok
}

func (e *Elev) takeAStep(p Point) ([]Point, error) {
	e.Visited[p] = true

	x := p.X
	y := p.Y

	candidates := []Point{
		{X: x + 1, Y: y}, // right
		{X: x, Y: y + 1}, // down
		{X: x - 1, Y: y}, // left
		{X: x, Y: y - 1}, // up
	}

	validCandidates := []Point{}
	for _, can := range candidates {
		if e.isValidStep(p, can) {
			validCandidates = append(validCandidates, can)
		}
	}
	if len(validCandidates) == 0 {
		return nil, errors.New("end of the line")
	}

	return validCandidates, nil
}

func (e *Elev) isValidStep(current, candidate Point) bool {
	if e.AlreadyVisited(candidate) {
		return false
	}

	maxX := len(e.SOS)
	maxY := len(e.SOS[0])
	if candidate.X > maxX-1 || candidate.X < 0 {
		return false
	}
	if candidate.Y > maxY-1 || candidate.Y < 0 {
		return false
	}
	if current == e.Start {
		return true
	}
	if candidate == e.Start {
		return false
	}
	if candidate == e.End { // Looks like you need to get to E via z
		return e.SOS[current.X][current.Y] == 122
	}

	currentVal := e.SOS[current.X][current.Y]
	newVal := e.SOS[candidate.X][candidate.Y]

	return newVal < currentVal+2 // Up one max (or down any)
}

//revive:disable:cognitive-complexity

func GetStepCount(elev Elev, startingPoint Point) int {
	startingPoints := []Point{startingPoint}
	for i := 0; i < 1000; i++ {
		newStartingPoints := []Point{}
		for _, start := range startingPoints {
			newPoints, err := elev.takeAStep(start) // No error the first time.
			if err != nil {
				continue
			}
			newStartingPoints = append(newStartingPoints, newPoints...)
		}

		// Deduplicate - and check for the end
		dumb := make(map[string]Point)
		for _, p := range newStartingPoints {
			if p == elev.End { // ARE WE DONE YET??
				return i + 1
			}
			dumb[fmt.Sprintf("%d,%d", p.X, p.Y)] = p
		}
		newNewStartingPoints := []Point{}
		for _, p := range dumb {
			newNewStartingPoints = append(newNewStartingPoints, p)
		}

		if len(newNewStartingPoints) == 0 {
			// panic("not expected to get here")//
			return 0
		}

		startingPoints = newNewStartingPoints
	}

	return 0 // Not expected
}

//revive:enable:cognitive-complexity
