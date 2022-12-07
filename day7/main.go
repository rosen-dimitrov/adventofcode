package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/foolishhumans/advent-of-code-lets-go/utils/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	rootDir := BuildTree(parseInput(input))
	return rootDir.ComputeNestedDirSize(100000)
}

func part2(input string) int {
	rootDir := BuildTree(parseInput(input))
	//init smallest to root
	smallest := rootDir.GetSize()

	need := 30000000 - (70000000 - smallest)
	rootDir.Walk(func(dir *Directory) {
		if dir.GetSize() >= need && dir.GetSize() <= smallest {
			smallest = dir.GetSize()
		}
	})

	return smallest
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}
