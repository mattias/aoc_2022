package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

func newFile(name string) *file {
	file := file{}
	file.name = name

	return &file
}

type directory struct {
	parentDirectory *directory
	name            string
	directories     []*directory
	files           []*file
}

func newDirectory(parentDirectory *directory, name string) *directory {
	dir := directory{}
	dir.parentDirectory = parentDirectory
	dir.name = name

	return &dir
}

func (dir *directory) isRoot() bool {
	return dir.parentDirectory == nil
}

func (dir *directory) size() int {
	size := 0

	for _, file := range dir.files {
		size += file.size
	}

	for _, directory := range dir.directories {
		size += directory.size()
	}

	return size
}

type filesystem struct {
	root directory
	cwd  *directory
}

func newFilesystem() *filesystem {
	root := directory{}
	root.name = "/"

	fs := filesystem{}
	fs.root = root
	fs.cwd = &root
	return &fs
}

func (fs *filesystem) cd(command string) {
	if command == ".." {
		if fs.cwd.isRoot() {
			return
		}

		fs.cwd = fs.cwd.parentDirectory
	}

	if command == "/" {
		fs.cwd = &fs.root
	}

	for _, dir := range fs.cwd.directories {
		if strings.TrimSpace(dir.name) == strings.TrimSpace(command) {
			fs.cwd = dir
			break
		}
	}
}

func (fs *filesystem) addDir(name string) {
	fs.cwd.directories = append(fs.cwd.directories, newDirectory(fs.cwd, strings.TrimSpace(name)))
}

func (fs *filesystem) addFile(size int, name string) {
	newFile := newFile(strings.TrimSpace(name))
	newFile.size = size
	fs.cwd.files = append(fs.cwd.files, newFile)
}

func (dir *directory) directorySizeArray() []int {
	sizes := []int{}

	for _, subDir := range dir.directories {
		sizes = append(sizes, subDir.directorySizeArray()...)
	}

	sizes = append(sizes, dir.size())
	return sizes
}

func getSolutionPart1(fs *filesystem) int {
	directorySizes := fs.root.directorySizeArray()

	totalSum := 0
	for _, size := range directorySizes {
		if size < 100000 {
			totalSum += size
		}
	}

	return totalSum
}

func getSolutionPart2(fs *filesystem) int {
	unusedSpace := 70000000 - fs.root.size()
	spaceToGet := 30000000 - unusedSpace

	directorySizes := fs.root.directorySizeArray()

	sort.Ints(directorySizes)

	for _, size := range directorySizes {
		if size > spaceToGet {
			return size
		}
	}

	return 0
}

func stringToInt(input string) (int, error) {
	i, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func parseInput(input string) *filesystem {
	fs := newFilesystem()

	lines := strings.Split(strings.TrimSpace(input), "\r\n")

	for _, line := range lines {
		if line[0] == '$' {
			if line[2:4] == "cd" {
				fs.cd(line[5:])
				continue
			}

			if line[2:4] == "ls" {
				continue
			}
		}

		if line[:3] == "dir" {
			fs.addDir(line[4:])
			continue
		}

		fileSplit := strings.Split(line, " ")
		size, _ := stringToInt(fileSplit[0])
		fs.addFile(size, fileSplit[1])
	}

	return fs
}

func main() {
	inputBytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic("couldn't read input")
	}

	input := parseInput(string(inputBytes))

	fmt.Println("Go")
	part := os.Getenv("part")

	if part == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}
