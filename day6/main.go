package main

import (
	"flag"
	"fmt"

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
	for i := 4; i < len(input); i++ {
		if hasMarker(input[i-4 : i]) {
			return i
		}
	}
	panic("im stupid if it gets to here")
}

func part2(input string) int {
	for i := 14; i < len(input); i++ {
		if hasMarker(input[i-14 : i]) {
			return i
		}
	}
	panic("im stupid if it gets to here")
}

func hasMarker(str string) bool {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}
	return true
}
