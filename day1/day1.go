package day1

import (
	"sort"
)

func Part1(data []int) int {
	elves := getElves(data)
	sort.Sort(ByFood(elves))

	return elves[0].total
}

func Part2(data []int) int {
	elves := getElves(data)
	sort.Sort(ByFood(elves))
	top3 := 0
	for i := 0; i < 3; i++ {
		top3 += elves[i].total
	}

	return top3
}

func getElves(data []int) []elf {
	elves := []elf{}
	index := 1
	e := elf{id: index}
	for _, cals := range data {
		if cals == 0 { // write the elf
			elves = append(elves, e)
			index++
			e = elf{id: index}
			continue
		}
		e.addFood(cals)
	}
	elves = append(elves, e) // write the final elf

	return elves
}

type elf struct {
	id    int
	food  []int
	total int
}

func (e *elf) addFood(f int) {
	e.food = append(e.food, f)
	e.total += f
}

type ByFood []elf

func (f ByFood) Len() int           { return len(f) }
func (f ByFood) Less(i, j int) bool { return f[i].total > f[j].total }
func (f ByFood) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
