package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/foolishhumans/advent-of-code-lets-go/utils/util"
	"k8s.io/utils/strings/slices"
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

	lines := parseInput(input)
	total := 0
	for _, line := range lines {
		total += calcualteWinnerForStandardGame(line)
	}
	return total
}

func part2(input string) int {
	lines := parseInput(input)
	total := 0
	for _, line := range lines {
		total += calcualteWinnerForSecretStrategy(line)
	}
	return total
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, "\n")
}

func findMove(enemyMove string, shouldWin bool) string {
	data := map[string]int{
		"A": 0, //rock
		"B": 1, //paper
		"C": 2, //scissors
	}
	enemyMoveBinary := data[enemyMove] | 1<<2

	for key := range data {

		myMoveBinary := data[key] | 0<<2
		if (enemyMoveBinary-myMoveBinary)%3 == 0 {
			if shouldWin {
				return key
			} else {
				wrong := []string{enemyMove, key}
				for k := range data {
					if !slices.Contains(wrong, k) {
						return k
					}
				}
			}
		}
	}
	panic("im dumb if it gets to here")
}

func calcualteWinnerForSecretStrategy(inputLine string) int {
	parsed := strings.Fields(inputLine)
	enemyMove := parsed[0]
	endResult := parsed[1]
	data := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
	}

	//draw
	if endResult == "Y" {
		return 3 + data[enemyMove] + 1
	}

	//I win
	if endResult == "Z" {
		return 6 + data[findMove(enemyMove, true)] + 1
	}
	//enemy wins
	return 0 + data[findMove(enemyMove, false)] + 1
}

// rock = 00 ; paper = 01; scissors = 10
// enemy = 1; me = 0
// 1st bit is the player; 2nd and 3rd are the choice
// rock/enemy = 00 -> 00 | 1 = 01 -> 01 << 2 -> 100 -> enemy rock = 100;
// paper/me = 01 -> 01 | 0 = 01 -> 01 << 2 -> 001 -> me paper = 001;
// enemy/rock - me/paper -> enemy=4(100bit); me=1(001bits) -> 4-1=3
// enemy/paper - me/scisors -> enemy=5(101bits); me=2(010) -> 5-2=3
// if enemyMove - myMove is multiple of 3 I win
func calcualteWinnerForStandardGame(inputLine string) int {
	move := strings.Fields(inputLine)
	data := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"X": 0,
		"Y": 1,
		"Z": 2,
	}
	//add 1 to bit to get the object value
	objectValue := data[move[1]] + 1

	//draw
	if data[move[0]] == data[move[1]] {
		return 3 + objectValue
	}

	//I win
	if ((data[move[0]]|1<<2)-
		(data[move[1]]|0<<2))%3 == 0 {
		return 6 + objectValue
	}

	//I lose
	return 0 + objectValue
}
