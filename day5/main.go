package main

import (
	"flag"
	"fmt"
	"strings"
	"unicode"

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

func part1(input string) string {

	stacks, commands := parseInput(input)
	var result string

	for _, command := range commands {
		for i := 1; i <= command[0]; i++ {
			stacks[command[2]-1] = string(stacks[command[1]-1][0]) + stacks[command[2]-1]
			stacks[command[1]-1] = stacks[command[1]-1][1:]
		}
	}

	for _, i := range stacks {
		result += string(i[0])
	}

	return result
}

func part2(input string) string {
	var result string

	stacks, commands := parseInput(input)
	for _, command := range commands {
		tmpStack := stacks[command[1]-1][:command[0]]
		stacks[command[1]-1] = stacks[command[1]-1][command[0]:]
		stacks[command[2]-1] = tmpStack + stacks[command[2]-1]
	}

	for _, i := range stacks {
		result += string(i[0])
	}

	return result
}

func parseInput(input string) (stacks []string, commands [][3]int) {
	stacksInput := strings.Split(strings.Split(string(input), "\n\n")[0], "\n")
	commandsInput := strings.Split(strings.Split(string(input), "\n\n")[1], "\n")
	stacks = parseStacks(stacksInput)
	commands = parseCommands(commandsInput)

	return stacks, commands
}

func parseCommand(s string) [3]int {
	var command [3]int
	fmt.Sscanf(s, "move %d from %d to %d", &command[0], &command[1], &command[2])

	return command
}

func parseColumn(crateInput []string, col int) string {
	var result strings.Builder

	for i := 0; i < len(crateInput); i++ {
		if crateLetter := crateInput[i][col]; unicode.IsLetter(rune(crateLetter)) {
			result.WriteByte(crateLetter)
		}
	}

	return result.String()
}

func parseStacks(stacksInput []string) (stacks []string) {
	for i := 1; i < len(stacksInput[0]); i += 4 {
		stacks = append(stacks, parseColumn(stacksInput, i))
	}
	return stacks
}

func parseCommands(commandsInput []string) (commands [][3]int) {
	for _, command := range commandsInput {
		commands = append(commands, parseCommand(command))
	}
	return commands
}
