package filesystem

import (
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

type Dir struct {
	Path      string
	ChildDirs []*Dir
	Files     []File // never change
}

type FileSystem struct {
	Directories map[string]*Dir
	Root        *Dir
	PWD         *Dir
}

type Command struct {
	Name string // cd, dir, filename
	Data string
}

func New() *FileSystem {
	root := &Dir{Path: "root"}
	directories := make(map[string]*Dir)
	directories["root"] = root
	fs := &FileSystem{
		Root:        root, // this shouldn't change. Change to rootname???
		PWD:         root,
		Directories: directories,
	}

	return fs
}

func (fs *FileSystem) CD(dirName string) *Dir {
	if dirName == "/" {
		return nil
	}
	newPath := getNewPath(fs.PWD.Path, dirName)
	newDir := fs.GetDirByPath(newPath)
	fs.PWD = newDir

	return newDir
}

func (fs *FileSystem) GetDirByPath(dirName string) *Dir {
	return fs.Directories[dirName]
}

func (fs *FileSystem) GetPWD() string {
	return fs.PWD.Path
}

func (fs *FileSystem) AddChildDirectory(dirName string) {
	newDir := fs.PWD.AddChildDirectory(dirName) // Always add to PWD
	fs.Directories[newDir.Path] = newDir        // add to registry
}

func (fs *FileSystem) ExecuteCommands(commands []Command) {
	for _, cmd := range commands {
		switch cmd.Name {
		case "cd":
			fs.CD(cmd.Data)
		case "dir":
			fs.AddChildDirectory(cmd.Data)
		default:
			fs.PWD.AddFile(cmd.Name, toInt(cmd.Data))
		}
	}
}

func (d *Dir) AddFile(filename string, size int) {
	d.Files = append(d.Files, File{Name: filename, Size: size})
}

// AddChildDirectory adds to the PWD
func (d *Dir) AddChildDirectory(path string) *Dir {
	childPath := getNewPath(d.Path, path)
	newChildDirectory := &Dir{Path: childPath}
	d.ChildDirs = append(d.ChildDirs, newChildDirectory)

	return newChildDirectory
}

func (d *Dir) getChildren(kids []*Dir) []*Dir {
	for _, child := range d.ChildDirs {
		kids = append(kids, child)
		kids = child.getChildren(kids)
	}

	return kids
}

func (d *Dir) GetFileSizeRecursive() int {
	allDirs := d.getChildren([]*Dir{d}) // seed with current dir
	total := 0
	for _, dir := range allDirs {
		for _, file := range dir.Files {
			total += file.Size
		}
	}

	return total
}

func getNewPath(pwdPath, path string) string {
	if path == ".." {
		parts := strings.Split(pwdPath, "/")
		parts = parts[:len(parts)-1] // drop the last element
		newPath := strings.Join(parts, "/")
		if newPath == "" {
			newPath = "root"
		}
		return newPath
	}

	return pwdPath + "/" + path
}

func toInt(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic("Not expected")
	}

	return val
}
