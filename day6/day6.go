package day6

func Day6(data string, offset int) int {
	for i := 0; i < len(data)-offset+1; i++ {
		chunk := data[i : i+offset]
		if allDiff(chunk) {
			return i + offset
		}
	}

	return 0 // Not expected
}

func allDiff(input string) bool {
	store := make(map[rune]bool)
	for _, run := range input {
		store[run] = true
	}

	return len(store) == len(input)
}
