package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/foolishhumans/advent-of-code-lets-go/utils/cast"
	"github.com/foolishhumans/advent-of-code-lets-go/utils/util"
)

type section struct {
	x int
	y int
}

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
	sections := parseInput(input)
	overlaps := 0
	for _, sectionPair := range sections {
		if fullOverlap(sectionPair[0], sectionPair[1]) || fullOverlap(sectionPair[1], sectionPair[0]) {
			overlaps++
		}
	}
	return overlaps
}

func part2(input string) int {
	sections := parseInput(input)
	overlaps := 0
	for _, sectionPair := range sections {
		if overlap(sectionPair[0], sectionPair[1]) {
			overlaps++
		}
	}
	return overlaps
}

func parseInput(input string) (ans [][]section) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		sectionIds := strings.FieldsFunc(line, func(r rune) bool {
			return r == '-' || r == ','
		})
		first := section{
			x: cast.ToInt(sectionIds[0]),
			y: cast.ToInt(sectionIds[1]),
		}
		second := section{
			x: cast.ToInt(sectionIds[2]),
			y: cast.ToInt(sectionIds[3]),
		}
		ans = append(ans, []section{first, second})
	}
	return ans
}

func fullOverlap(first, second section) bool {
	return second.x >= first.x && second.x <= first.y && second.y >= first.x && second.y <= first.y
}

func overlap(first, second section) bool {
	return !(first.x > second.y || first.y < second.x)
}
