package grid

func New() *Grid {
	start := Point{0, 0}
	visited := make(map[Point]bool)
	visited[start] = true

	return &Grid{Visited: visited, Head: start, Tail: start}
}

type Grid struct {
	Head    Point
	Tail    Point
	Visited map[Point]bool
}

func (t *Grid) Move(dir string, moveCount int) {
	switch dir {
	case "L":
		t.moveLeft(moveCount)
	case "R":
		t.moveRight(moveCount)
	case "U":
		t.moveUp(moveCount)
	case "D":
		t.moveDown(moveCount)
	}
}

func isClose(head, tail Point) bool {
	xDiff := head.X - tail.X
	xDiff = abs(xDiff)
	yDiff := head.Y - tail.Y
	yDiff = abs(yDiff)

	return xDiff < 2 && yDiff < 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (t *Grid) moveRight(moveCount int) []Point {
	t.Head = Point{t.Head.X + moveCount, t.Head.Y}
	if isClose(t.Head, t.Tail) {
		return nil
	}
	visitedThisMove := []Point{} // for debug
	i := 1
	for {
		step := Point{t.Tail.X + i, t.Head.Y}
		visitedThisMove = append(visitedThisMove, step)
		t.Visited[step] = true
		if isClose(t.Head, step) {
			break
		}
		i++
	}
	t.Tail = visitedThisMove[len(visitedThisMove)-1]

	return visitedThisMove
}

func (t *Grid) moveLeft(moveCount int) []Point {
	t.Head = Point{t.Head.X - moveCount, t.Head.Y}
	if isClose(t.Head, t.Tail) {
		return nil
	}
	visitedThisMove := []Point{} // for debug
	i := 1
	for {
		step := Point{t.Tail.X - i, t.Head.Y}
		visitedThisMove = append(visitedThisMove, step)
		t.Visited[step] = true
		if isClose(t.Head, step) {
			break
		}
		i++
	}
	t.Tail = visitedThisMove[len(visitedThisMove)-1]

	return visitedThisMove
}

func (t *Grid) moveUp(moveCount int) []Point {
	t.Head = Point{t.Head.X, t.Head.Y + moveCount}
	if isClose(t.Head, t.Tail) {
		return nil
	}
	visitedThisMove := []Point{} // for debug
	i := 1
	for {
		step := Point{t.Head.X, t.Tail.Y + i}
		visitedThisMove = append(visitedThisMove, step)
		t.Visited[step] = true
		if isClose(t.Head, step) {
			break
		}
		i++
	}
	t.Tail = visitedThisMove[len(visitedThisMove)-1]

	return visitedThisMove
}

func (t *Grid) moveDown(moveCount int) []Point {
	t.Head = Point{t.Head.X, t.Head.Y - moveCount}
	if isClose(t.Head, t.Tail) {
		return nil
	}
	visitedThisMove := []Point{} // for debug
	i := 1
	for {
		step := Point{t.Head.X, t.Tail.Y - i}
		visitedThisMove = append(visitedThisMove, step)
		t.Visited[step] = true
		if isClose(t.Head, step) {
			break
		}
		i++
	}
	t.Tail = visitedThisMove[len(visitedThisMove)-1]

	return visitedThisMove
}

type Point struct {
	X int
	Y int
}

// Plot is used for debugging only
// func (t *Grid) Plot() {
// 	limit := 5 // hacking
// 	for y := limit; y > -1; y-- {
// 		fmt.Println("")
// 		for x := 0; x < limit; x++ {
// 			_, ok := t.Visited[Point{X: x, Y: y}]
// 			if ok {
// 				fmt.Print("X")
// 			} else {
// 				fmt.Print(".")
// 			}
// 		}
// 	}
// }
