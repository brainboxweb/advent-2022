package troop

import (
	"sort"
	"strconv"
)

type Troop struct {
	Monkeys      []*Monkey
	worryDivisor int
}

func New(worryDivisor int) *Troop {
	return &Troop{worryDivisor: worryDivisor}
}

func (tt *Troop) AddMonkey(
	id int,
	items []int,
	operator,
	operand string,
	test,
	throwTrue,
	throwFalse int,
) {
	m := &Monkey{
		ID:         id,
		Items:      items,
		Operator:   operator,
		Operand:    operand,
		Test:       test,
		ThrowTrue:  throwTrue,
		ThrowFalse: throwFalse,
	}
	tt.Monkeys = append(tt.Monkeys, m)
}

func (tt *Troop) Catch(monkeyID, item int) {
	tt.Monkeys[monkeyID].AddItem(item)
}

func (tt *Troop) MonkeyBusiness() int {
	insp := []int{}
	for _, m := range tt.Monkeys {
		insp = append(insp, m.InspectionCount)
	}
	sort.Ints(insp)
	score := insp[len(insp)-1] * insp[len(insp)-2]

	return score
}

type Monkey struct {
	ID              int
	Items           []int
	Operator        string
	Operand         string
	ThrowTrue       int
	ThrowFalse      int
	Test            int
	InspectionCount int
}

func (m *Monkey) AddItem(item int) {
	m.Items = append(m.Items, item)
}

type Throw struct {
	MonkeyID int
	Item     int
}

func (m *Monkey) Inspect(worryDivisor int) []Throw {
	throws := []Throw{}
	for _, item := range m.Items { // rename this
		m.InspectionCount++
		operand := getOperand(m.Operand, item)
		newWorry := 0
		switch m.Operator {
		case "*":
			newWorry = operand * item
		case "+":
			newWorry = operand + item
		default:
			panic("Not expected")
		}
		boredLevel := newWorry / worryDivisor
		if boredLevel%m.Test == 0 {
			t := Throw{m.ThrowTrue, boredLevel}
			throws = append(throws, t)
		} else {
			t := Throw{m.ThrowFalse, boredLevel}
			throws = append(throws, t)
		}
	}
	// remove the items (assume all get thrown)
	m.Items = nil

	return throws
}

func getOperand(monkeyOperand string, item int) int {
	if monkeyOperand == "old" {
		return item
	}
	val, err := strconv.Atoi(monkeyOperand)
	if err != nil {
		panic("Not expected")
	}

	return val
}

func (tt *Troop) PlayRounds(roundCount int) {
	for i := 0; i < roundCount; i++ {
		for _, m := range tt.Monkeys {
			throws := m.Inspect(tt.worryDivisor)
			for _, throw := range throws {
				tt.Catch(throw.MonkeyID, throw.Item)
			}
		}
	}
}
