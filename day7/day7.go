package day7

import (
	"sort"
	"strings"

	"github.com/brainboxweb/advent-2022/day7/filesystem"
)

func Part1(data []string) int {
	cmds := parseData(data)

	fs := filesystem.New()
	fs.ExecuteCommands(cmds)
	allTheDirs := fs.Directories
	ret := 0
	for _, dir := range allTheDirs {
		totFileSize := dir.GetFileSizeRecursive()
		if totFileSize <= 100000 {
			ret += totFileSize
		}
	}

	return ret
}

func Part2(data []string) int {
	cmds := parseData(data)
	fs := filesystem.New()
	fs.ExecuteCommands(cmds)

	allTheDirs := fs.Directories
	sizes := make(map[string]int)
	for _, dir := range allTheDirs {
		totFileSize := dir.GetFileSizeRecursive()
		sizes[dir.Path] = totFileSize
	}

	freeSpace := 70000000 - sizes["root"]
	needed := 30000000 - freeSpace
	candidates := []int{}
	for _, size := range sizes {
		if size >= needed {
			candidates = append(candidates, size)
		}
	}
	sort.Ints(candidates)

	return candidates[0]
}

func parseData(data []string) []filesystem.Command {
	commands := []filesystem.Command{}
	for _, item := range data {
		var cmd filesystem.Command
		parts := strings.Split(item, " ")
		switch parts[0] {
		case "$":
			if parts[1] == "cd" {
				cmd = filesystem.Command{Name: "cd", Data: parts[2]}
				commands = append(commands, cmd)
			}
		case "dir":
			cmd = filesystem.Command{Name: "dir", Data: parts[1]}
			commands = append(commands, cmd)

		default: // file
			cmd = filesystem.Command{Name: parts[1], Data: parts[0]}
			commands = append(commands, cmd)
		}
	}

	return commands
}
