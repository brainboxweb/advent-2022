package day4

import (
	"strconv"
	"strings"
)

func Part1(data []string) int {
	ret := 0
	for _, item := range data {
		a, b := getParts(item)
		if a[0] <= b[0] && a[len(a)-1] >= b[len(b)-1] {
			ret++
			continue
		}
		if b[0] <= a[0] && b[len(b)-1] >= a[len(a)-1] {
			ret++
		}
	}

	return ret
}

func Part2(data []string) int {
	ret := 0
	for _, item := range data {
		a, b := getParts(item)
		switch {
		case a[len(a)-1] >= b[0] && b[len(b)-1] >= a[0]:
			ret++
			continue
		case b[len(b)-1] >= a[0] && a[len(a)-1] >= b[0]:
			ret++
			continue
		}
	}

	return ret
}

func getParts(input string) (a, b []int) {
	parts := strings.Split(input, ",")

	aRange := strings.Split(parts[0], "-")
	bRange := strings.Split(parts[1], "-")

	aL, _ := strconv.Atoi(aRange[0])
	aH, _ := strconv.Atoi(aRange[1])

	bL, _ := strconv.Atoi(bRange[0])
	bH, _ := strconv.Atoi(bRange[1])

	a = []int{aL, aH}
	b = []int{bL, bH}

	return a, b
}
