package day2

import "strings"

func Part1(data []string) int {
	score := 0
	for _, line := range data {
		parts := strings.Split(line, " ")
		oppo := getRPS(parts[0])
		me := getRPS(parts[1])
		result := turn(oppo, me)
		score += result
	}

	return score
}

func Part2(data []string) int {
	score := 0
	for _, line := range data {
		parts := strings.Split(line, " ")
		oppo := getRPS(parts[0])
		me := getMove(oppo, parts[1])
		result := turn(oppo, me)
		score += result
	}

	return score
}

func getRPS(in string) string {
	switch in {
	case "A":
		return "rock"
	case "B":
		return "paper"
	case "C":
		return "scissors"
	case "X":
		return "rock"
	case "Y":
		return "paper"
	case "Z":
		return "scissors"
	}
	panic("Not expected")
}

func getMove(oppo, required string) string {
	/*
		X means you need to lose,
		Y means you need to end the round in a draw, and
		Z means you need to win. Good luck!"
	*/
	switch required {
	case "X": // lose
		switch oppo {
		case "rock": // Oppo plays...
			return "scissors"
		case "paper":
			return "rock"
		case "scissors":
			return "paper"
		}
	case "Y": // draw
		return oppo
	case "Z": // win
		switch oppo {
		case "rock": // Oppo plays...
			return "paper"
		case "paper":
			return "scissors"
		case "scissors":
			return "rock"
		}
	}

	panic("Not expected")
}

func turn(oppo, me string) int { // input is rock, paper or scissors
	items := make(map[string]int)
	items["rock"] = 1
	items["paper"] = 2
	items["scissors"] = 3
	ret := items[me]
	if oppo == me {
		ret += 3
		return ret
	}
	if me == "rock" && oppo == "scissors" ||
		me == "paper" && oppo == "rock" ||
		me == "scissors" && oppo == "paper" {
		ret += 6
	}

	return ret
}
