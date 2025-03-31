package day3

func Part1(data []string) int {
	total := 0
	for _, sack := range data {
		one := sack[:len(sack)/2]
		two := sack[len(sack)/2:]
		common := getCommon(one, two)
		total += runeToValue(common)
	}

	return total
}

func Part2(data []string) int {
	total := 0
	i := 0
	for {
		if i > len(data)-3 {
			break
		}
		common := getCommonFromThree(data[i], data[i+1], data[i+2])
		val := runeToValue(common)
		total += val
		i += 3
	}

	return total
}

func getCommon(one, two string) rune {
	for _, o := range one {
		for _, t := range two {
			if o == t {
				return t
			}
		}
	}
	panic("not expected")
}

//revive:disable:cognitive-complexity
func getCommonFromThree(one, two, three string) rune {
	for _, on := range one {
		for _, tw := range two {
			for _, tt := range three {
				if tw == on && tt == on {
					return on
				}
			}
		}
	}
	panic("not expected")
}

//revive:enable:cognitive-complexity

func runeToValue(r rune) int {
	val := int(r)
	if val < 97 { // uppercase chars
		return val - 64 + 26
	}
	return int(r) - 96 // lowercase chars
}
