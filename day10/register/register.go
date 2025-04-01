package register

import (
	"strconv"
	"strings"
)

func New() *Register {
	reg := make(map[int]int)
	reg[0] = 1

	return &Register{reg: reg}
}

type Register struct {
	cycle     int
	reg       map[int]int
	waitValue int
}

func (r *Register) increment() {
	currentCycle := r.cycle
	r.cycle++
	r.reg[currentCycle+1] = r.reg[currentCycle] // store the value
}

func (r *Register) GetStrength(index int) int {
	return r.reg[index] * index
}

func (r *Register) Len() int {
	return len(r.reg)
}

func (r *Register) GetValue(index int) int {
	return r.reg[index]
}

func (r *Register) ExecCommand(cmd string) {
	r.increment()
	if r.waitValue != 0 {
		r.reg[r.cycle] += r.waitValue
	}
	if cmd == "noop" { // Just increments the cycle (value doesn't change)
		r.waitValue = 0
		return
	}
	//
	r.increment()
	parts := strings.Split(cmd, " ")
	r.waitValue, _ = strconv.Atoi(parts[1]) // error handling?
}
