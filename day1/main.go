package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/foolishhumans/advent-of-code-lets-go/utils/cast"
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
	return max(input)
}

func part2(input string) int {
	var currentCal int = 0
	var totals []int
	lines := strings.Split(input, "\n")

	for i, num := range lines {
		if num != "" {
			currentCal += cast.ToInt(num)
		} else {
			totals = append(totals, currentCal)
			currentCal = 0
		}

		//cover a case where last elements dont have an empty line at the end
		if i == len(lines)-1 && currentCal > 0 {
			totals = append(totals, currentCal)
		}
	}

	sort.Ints(totals)
	return sum([]int{totals[len(totals)-1], totals[len(totals)-2], totals[len(totals)-3]})
}

func max(input string) int {
	lines := strings.Split(input, "\n")

	var max int = 0
	var currentCal int = 0
	for _, num := range lines {
		if num != "" {
			currentCal += cast.ToInt(num)
		} else {
			currentCal = 0
		}
		if currentCal >= max {
			max = currentCal
		}
	}
	return max
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
