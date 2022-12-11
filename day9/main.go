package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/foolishhumans/advent-of-code-lets-go/utils/cast"
	"github.com/foolishhumans/advent-of-code-lets-go/utils/util"
)

const (
	up    = "U"
	right = "R"
	down  = "D"
	left  = "L"
)

type point struct {
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
	return simulate(input, 2)
}

func part2(input string) int {
	return simulate(input, 10)
}

func (p *point) moveTowards(other point) {
	dx := util.Abs(other.x - p.x)
	dy := util.Abs(other.y - p.y)

	switch {
	case dx >= 2 && dy == 0:
		if other.x > p.x {
			// R
			p.x++
		} else {
			// L
			p.x--
		}
	case dy >= 2 && dx == 0:
		if other.y > p.y {
			// U
			p.y++
		} else {
			// D
			p.y--
		}
	case dx >= 2 && dy >= 1 || dy >= 2 && dx >= 1:
		// Diagonal
		if other.x > p.x {
			p.x++
			if other.y > p.y {
				p.y++
			} else {
				p.y--
			}
		} else {
			p.x--
			if other.y > p.y {
				p.y++
			} else {
				p.y--
			}
		}
	}
}

func simulate(input string, size int) int {
	lines := strings.Split(input, "\n")
	rope := make([]point, size)

	tailSeen := map[point]struct{}{
		{}: {},
	}

	for _, move := range lines {
		parts := strings.Fields(move)

		dir := parts[0]
		mag := cast.ToInt(parts[1])

		for i := 0; i < mag; i++ {
			switch dir {
			case up:
				rope[0].y++
			case right:
				rope[0].x++
			case down:
				rope[0].y--
			case left:
				rope[0].x--
			}

			for n := 1; n < len(rope); n++ {
				rope[n].moveTowards(rope[n-1])
			}

			tailSeen[rope[len(rope)-1]] = struct{}{}
		}
	}

	return len(tailSeen)
}
