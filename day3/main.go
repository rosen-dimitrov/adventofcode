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
	total := 0
	backpacks := parseInput(input)

	for _, backpack := range backpacks {
		left, right := partitionBackpack(backpack)
		total += calcultePriority(left & right)
	}
	return total
}

func part2(input string) int {
	total := 0

	unparsed := parseInput(input)
	for i := 0; i < len(unparsed); i += 3 {
		shared := pack(unparsed[i])
		shared &= pack(unparsed[i+1])
		shared &= pack(unparsed[i+2])

		total += calcultePriority(shared)
	}

	return total
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, "\n")
}

func partitionBackpack(backpack string) (left, right uint64) {
	return pack(backpack[:len(backpack)/2]), pack(backpack[len(backpack)/2:])
}

// use bitmap to store the index for each rune
func pack(backpack string) uint64 {
	var itemMap uint64
	for _, item := range backpack {
		if item >= 'a' && item <= 'z' {
			itemMap |= 1 << (item - 'a')
		} else {
			itemMap |= 1 << (item - 'A' + 26)
		}
	}

	return itemMap
}

func calcultePriority(backpack uint64) int {
	var total int

	for i := 0; i < 26*2; i++ {
		//check if there is a flag in the bitmap for this item
		//0b1 == 1 in decimal; backpack consists of runes
		if (backpack>>i)&0b1 == 1 {
			total += i + 1
		}
	}

	return total
}
