package main

import (
	"flag"
	"fmt"
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
	matrix, cols := parseInput(input)

	visible := 0
	for ri, row := range matrix {
		for ti, tree := range row {
			// top, bottom, left and right are always visible
			if ri == 0 || ti == 0 || ri == len(matrix)-1 || ti == len(row)-1 {
				visible++
				// right
			} else if treeIsVisible(row[ti+1:], tree) {
				visible++
				// left
			} else if treeIsVisible(row[:ti], tree) {
				visible++
				// below
			} else if treeIsVisible(cols[ti][ri+1:], tree) {
				visible++
				// above
			} else if treeIsVisible(cols[ti][:ri], tree) {
				visible++
			}
		}
	}
	return visible
}

func part2(input string) int {
	matrix, cols := parseInput(input)
	max := 0
	for ri, row := range matrix {
		for ti, tree := range row {
			right := visibleTrees(row[ti+1:], tree, false)      // right
			left := visibleTrees(row[:ti], tree, true)          // left
			below := visibleTrees(cols[ti][ri+1:], tree, false) // below
			above := visibleTrees(cols[ti][:ri], tree, true)    // above
			score := right * left * below * above

			if score > max {
				max = score
			}
		}
	}
	return max
}

func parseInput(input string) ([][]int, [][]int) {
	lines := strings.Split(input, "\n")
	matrix := buildMatrix(lines)
	columns := getColumns(matrix)

	return matrix, columns
}

func buildMatrix(input []string) [][]int {
	matrix := make([][]int, len(input))

	for i, line := range input {
		trees := strings.Split(line, "")
		matrix[i] = make([]int, 0)
		for _, tree := range trees {
			treeHight := cast.ToInt(tree)
			matrix[i] = append(matrix[i], treeHight)
		}
	}
	return matrix
}

func getColumns(matrix [][]int) [][]int {
	columns := make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		columns[i] = make([]int, 0)
		for j := 0; j < len(matrix); j++ {
			columns[i] = append(columns[i], matrix[j][i])
		}
	}
	return columns
}

func treeIsVisible(row []int, n int) bool {
	for _, v := range row {
		if v >= n {
			return false
		}
	}
	return true
}

func visibleTrees(row []int, n int, reverse bool) int {
	count := 0
	if reverse {
		for i := len(row) - 1; i >= 0; i-- {
			count++
			if row[i] >= n {
				return count
			}
		}
	} else {
		for _, v := range row {
			count++
			if v >= n {
				return count
			}
		}
	}

	return count
}
