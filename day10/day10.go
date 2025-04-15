package day10

import (
	"strings"

	"github.com/brainboxweb/advent-2022/day10/register"
)

func Part1(data []string) int {
	reg := register.New()
	for _, cmd := range data {
		reg.ExecCommand(cmd)
	}
	factors := []int{20, 60, 100, 140, 180, 220}
	ret := 0
	for _, factor := range factors {
		ret += reg.GetStrength(factor)
	}

	return ret
}

func Part2(data []string) string {
	var builder strings.Builder
	reg := register.New()
	for _, cmd := range data {
		reg.ExecCommand(cmd)
	}
	for i := 1; i < reg.Len(); i++ {
		val := reg.GetValue(i)
		sprite := i % 40
		if sprite == 0 {
			_, _ = builder.WriteString("\n")
			continue
		}
		if sprite == val || sprite-1 == val || sprite-2 == val {
			_, _ = builder.WriteString("#")
			continue
		}
		_, _ = builder.WriteString(".")
	}
	output := builder.String()

	return strings.Trim(output, "\n")
}
