package filesystem_test

import (
	"testing"

	"github.com/brainboxweb/advent-2022/day7/filesystem"
	"github.com/stretchr/testify/assert"
)

func TestExecuteCommands(t *testing.T) {
	tests := []struct {
		commands []filesystem.Command
		expected string
	}{
		{
			[]filesystem.Command{
				{"dir", "a"},
				{"cd", "a"},
				{"dir", "x"},
				{"dir", "y"},
				{"dir", "z"},
				{"cd", "z"},
				{"error.log", "1234567"},
			},
			"root/a/z",
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			fs := filesystem.New()
			fs.ExecuteCommands(tt.commands)
			assert.Equal(t, tt.expected, fs.GetPWD())
		})
	}
}

func TestGetDirByPath(t *testing.T) {
	fs := filesystem.New()
	fs.AddChildDirectory("a")
	fs.CD("a")
	fs.AddChildDirectory("b")
	fs.CD("b")
	fs.AddChildDirectory("c")
	fs.CD("c")

	assert.Equal(t, "root/a/b/c", fs.GetPWD())
}

func TestCD(t *testing.T) {
	fs := filesystem.New()
	assert.Equal(t, "root", fs.GetPWD())

	fs.AddChildDirectory("a")
	fs.CD("a")
	assert.Equal(t, "root/a", fs.GetPWD())

	fs.CD("..")
	assert.Equal(t, "root", fs.GetPWD())

	fs.CD("a")
	assert.Equal(t, "root/a", fs.GetPWD())

	fs.AddChildDirectory("b")
	fs.CD("b")
	assert.Equal(t, "root/a/b", fs.GetPWD())

	fs.CD("..")
	assert.Equal(t, "root/a", fs.GetPWD())

	fs.CD("..")
	assert.Equal(t, "root", fs.GetPWD())

	fs.CD("a")
	assert.Equal(t, "root/a", fs.GetPWD())

	fs.CD("..")
	assert.Equal(t, "root", fs.GetPWD())

	fs.AddChildDirectory("c")
	fs.CD("c")
	assert.Equal(t, "root/c", fs.GetPWD())

	fs.AddChildDirectory("d")
	fs.CD("d")
	assert.Equal(t, "root/c/d", fs.GetPWD())
}

func TestGetFileSizeRecursive(t *testing.T) {
	fs := filesystem.New()
	fs.AddChildDirectory("a")

	dirA := fs.CD("a")
	dirA.AddFile("hello", 1000)

	fs.AddChildDirectory("b")
	dirB := fs.CD("b")
	dirB.AddFile("hello", 1000000)

	totSize := dirA.GetFileSizeRecursive()

	assert.Equal(t, 1001000, totSize)
}
