package helpers

import (
	"bufio"
	"os"
	"strconv"
)

func ReverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func ReverseSliceOfSlices(s [][]string) [][]string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func GetData(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic("not expected")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var dat []int
	for scanner.Scan() {
		if scanner.Text() == "" {
			dat = append(dat, 0)
			continue
		}
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("not expected")
		}
		dat = append(dat, i)
	}
	return dat
}

func GetDataString(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic("not expected")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var dat []string
	for scanner.Scan() {
		dat = append(dat, scanner.Text())
	}
	return dat
}
