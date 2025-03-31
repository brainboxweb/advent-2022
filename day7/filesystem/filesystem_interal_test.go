package filesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddChildDirectory(t *testing.T) {
	fs := New()
	childDir := fs.PWD.AddChildDirectory("a")
	assert.Equal(t, "root/a", childDir.Path)
}

func TestNewPath(t *testing.T) {
	tests := []struct {
		pwdPath  string
		path     string
		expected string
	}{
		{
			"root",
			"a",
			"root/a",
		},
		{
			"root/a",
			"..",
			"root",
		},
		{
			"root",
			"..",
			"root",
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result := getNewPath(tt.pwdPath, tt.path)
			assert.Equal(t, tt.expected, result)
		})
	}
}
