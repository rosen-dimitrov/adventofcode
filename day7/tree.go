package main

import (
	"strings"

	"github.com/foolishhumans/advent-of-code-lets-go/utils/cast"
)

// files are directories with no children
// store tree as map of dirpath and dirpath
type Directory struct {
	children        map[string]*Directory
	filesSize       int
	nestedFilesSize int
}

func BuildTree(lines []string) *Directory {
	root := &Directory{children: map[string]*Directory{}}

	var dirpath []string
	for _, line := range lines {
		lineElements := strings.Fields(line)
		//if it's cd command append the new dir to the dirpath
		if isCdCommand(lineElements) {
			switch lineElements[2] {
			case "..":
				dirpath = dirpath[:len(dirpath)-1]
			case "/":
				dirpath = nil
			default:
				dirpath = append(dirpath, lineElements[2])
			}

			continue
		}

		// ignore ls command and dir line
		if lineElements[0] == "dir" || isLsCommand(lineElements) {
			continue
		}

		//pointer to currentdir
		current := root
		//recursively init struct tree
		for _, path := range dirpath {
			if _, ok := current.children[path]; !ok {
				current.children[path] = &Directory{children: map[string]*Directory{}}
			}
			current = current.children[path]
		}
		//add current lines' file size to the current directory's total
		current.filesSize += cast.ToInt(lineElements[0])
	}

	return root
}

// just ignore the ls command as it's not needed for anything
func isCdCommand(line []string) bool {
	return line[0] == "$" && line[1] == "cd"
}

func isLsCommand(line []string) bool {
	return line[0] == "$" && line[1] == "ls"
}

// recursively traverse tree and compute size for each subdirectory
func (dir *Directory) GetSize() int {
	if dir.nestedFilesSize != 0 {
		return dir.nestedFilesSize
	}

	total := dir.filesSize
	for _, child := range dir.children {
		total += child.GetSize()
	}

	dir.nestedFilesSize = total
	return total
}

// recursively calculate nested dir size
func (dir *Directory) ComputeNestedDirSize(maxSize int) int {
	var total int
	if size := dir.GetSize(); size <= maxSize {
		total += size
	}

	for _, child := range dir.children {
		total += child.ComputeNestedDirSize(maxSize)
	}

	return total
}

// recursive callback
func (dir *Directory) Walk(walker func(*Directory)) {
	walker(dir)
	for _, child := range dir.children {
		child.Walk(walker)
	}
}
