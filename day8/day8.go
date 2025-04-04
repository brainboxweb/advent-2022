package day8

import (
	"strconv"
)

func Part1(data []string) int {
	limit := len(data) // assume square
	mapData := parse(data)

	sliceOfVisibilities := make([][]bool, limit)
	for i := 0; i < limit; i++ {
		sliceOfVisibilities[i] = make([]bool, limit)
	}

	// Check every direction
	right(limit, mapData, sliceOfVisibilities)
	left(limit, mapData, sliceOfVisibilities)
	up(limit, mapData, sliceOfVisibilities)
	down(limit, mapData, sliceOfVisibilities)
	ret := 0
	for x := 0; x < limit; x++ {
		for y := 0; y < limit; y++ {
			if sliceOfVisibilities[x][y] {
				ret++
			}
		}
	}

	return ret
}

func down(limit int, mapData [][]int, sliceOfVisibilities [][]bool) {
	for x := 0; x < limit; x++ {
		hi := -1
		for y := limit - 1; y > -1; y-- {
			val := mapData[x][y]
			if val > hi { // visible!
				sliceOfVisibilities[x][y] = true
				hi = val
			}
		}
	}
}

func up(limit int, mapData [][]int, sliceOfVisibilities [][]bool) {
	for x := 0; x < limit; x++ {
		hi := -1
		for y := 0; y < limit; y++ {
			val := mapData[x][y]
			if val > hi { // visible!
				sliceOfVisibilities[x][y] = true
				hi = val
			}
		}
	}
}

func left(limit int, mapData [][]int, sliceOfVisibilities [][]bool) {
	for y := 0; y < limit; y++ {
		hi := -1
		for x := limit - 1; x > -1; x-- {
			val := mapData[x][y]
			if val > hi { // visible!
				sliceOfVisibilities[x][y] = true
				hi = val
			}
		}
	}
}

func right(limit int, mapData [][]int, sliceOfVisibilities [][]bool) {
	for y := 0; y < limit; y++ {
		hi := -1
		for x := 0; x < limit; x++ {
			val := mapData[x][y]
			if val > hi { // visible!
				sliceOfVisibilities[x][y] = true
				hi = val
			}
		}
	}
}

func Part2(data []string) int {
	limit := len(data) // assume square
	mapData := parse(data)
	max := 0
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			val := check(mapData, i, j)
			if val > max {
				max = val
			}
		}
	}

	return max
}

func parse(data []string) [][]int {
	limit := len(data) // assume square
	ret := make([][]int, limit)
	for i := 0; i < limit; i++ {
		ret[i] = make([]int, limit)
		line := data[i]
		lineItems := []int{}
		for _, theChar := range line {
			charInt, _ := strconv.Atoi(string(theChar))
			lineItems = append(lineItems, charInt)
		}
		ret[i] = lineItems
	}

	return ret
}

func check(input [][]int, i, j int) int {
	limit := len(input[0])
	current := input[i][j]
	count := countLeft(i, input, j, current)
	count *= countRight(i, limit, input, j, current)
	count *= countUp(j, limit, input, i, current)
	count *= countDown(j, input, i, current)

	return count
}

func countDown(j int, input [][]int, i int, current int) int {
	count := 0
	if j > 0 {
		for y := j - 1; y > -1; y-- {
			if input[i][y] >= current {
				count++
				break
			}
			count++
		}
	}
	return count
}

func countUp(j int, limit int, input [][]int, i int, current int) int {
	count := 0
	if j < limit {
		for y := j + 1; y < limit; y++ {
			if input[i][y] >= current {
				count++
				break
			}
			count++
		}
	}
	return count
}

func countLeft(i int, input [][]int, j int, current int) int {
	count := 0
	if i > 0 {
		for x := i - 1; x > -1; x-- {
			if input[x][j] >= current {
				count++
				break
			}
			count++
		}
	}
	return count
}

func countRight(i int, limit int, input [][]int, j int, current int) int {
	count := 0
	if i < limit {
		for x := i + 1; x < limit; x++ {
			if input[x][j] >= current {
				count++
				break
			}
			count++
		}
	}
	return count
}
