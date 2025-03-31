package day11

import (
	"sort"
	"strconv"
	"strings"

	"github.com/brainboxweb/advent-2022/day11/troop"
)

func Part1(data []string) int {
	rounds := 20
	worryDivisor := 3
	mmData := parse(data)
	tt := troop.New(worryDivisor)
	for _, mData := range mmData {
		tt.AddMonkey(
			mData.ID,
			mData.Items,
			mData.Operator,
			mData.Operand,
			mData.Test,
			mData.ThrowTrue,
			mData.ThrowFalse,
		)
	}
	tt.PlayRounds(rounds)

	return tt.MonkeyBusiness()
}

type MonkeyData struct {
	ID         int
	Items      []int
	Operator   string
	Operand    string
	Test       int
	ThrowTrue  int
	ThrowFalse int
}

//revive:disable:cyclomatic
//revive:disable:cognitive-complexity

func parse(data []string) []MonkeyData {
	monkeyIndex := 0
	ms := make(map[int]MonkeyData)
	for _, line := range data {
		parts := strings.Split(line, " ")
		switch {
		case line == "":
			monkeyIndex++
		case parts[0] != "":
			m := MonkeyData{ID: monkeyIndex}
			ms[monkeyIndex] = m
		case parts[2] != "":
			switch parts[2] {
			case "Starting":
				m := ms[monkeyIndex]
				for i := 4; i < len(parts); i++ {
					val, err := strconv.Atoi(strings.Trim(parts[i], ","))
					if err != nil {
						panic("not expected")
					}
					m.Items = append(m.Items, val)
				}
				ms[monkeyIndex] = m
			case "Operation:":
				m := ms[monkeyIndex]
				m.Operator = parts[6]
				m.Operand = parts[7]
				ms[monkeyIndex] = m
			case "Test:":
				m := ms[monkeyIndex]
				val, err := strconv.Atoi(parts[5])
				if err != nil {
					panic("not expected")
				}
				m.Test = val
				ms[monkeyIndex] = m
			}
		case parts[4] != "":
			mIndex, err := strconv.Atoi(parts[9])
			if err != nil {
				panic("not expected")
			}
			m := ms[monkeyIndex]
			if parts[5] == "true:" {
				m.ThrowTrue = mIndex
			} else {
				m.ThrowFalse = mIndex
			}
			ms[monkeyIndex] = m
		}
	}
	ret := sortTroopData(ms) // !important!

	return ret
}

//revive:enable:cyclomatic
//revive:enable:cognitive-complexity

func sortTroopData(m map[int]MonkeyData) []MonkeyData {
	ret := []MonkeyData{}
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		ret = append(ret, m[k])
	}

	return ret
}
