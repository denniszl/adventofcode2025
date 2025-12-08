package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput() [][]string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	splitted := strings.Split(string(dat), "\n")
	splitted = splitted[0 : len(splitted)-1]
	out := [][]string{}
	for _, s := range splitted {
		a := []string{}
		for _, c := range s {
			a = append(a, string(c))
		}
		out = append(out, a)
	}
	return out
}

func printDiagram(input [][]string) {
	out := []string{}
	for _, v := range input {
		out = append(out, strings.Join(v, ""))
	}

	fmt.Println(strings.Join(out, "\n"))
}

func drawBeam(x, y int, input [][]string) {
	if y == len(input) {
		return
	}
	if input[y][x] == "^" {
		return
	}
	input[y][x] = "|"
}

var outputPositions = map[string]bool{}
var visited = map[string]bool{}
var memo = map[string]int{}

func calculateBeam(x, y int, input [][]string) int {
	if x < 0 || x >= len(input[0]) {
		return 0
	}
	posKey := fmt.Sprintf("%d,%d", y, x)
	if _, ok := memo[posKey]; ok {
		return memo[posKey]
	}
	if y >= len(input) {
		memo[posKey] = 1
		return 1
	}

	// posKey := fmt.Sprintf("%d,%d", y, x)
	// if visited[posKey] {
	// 	return 0
	// }
	// visited[posKey] = true

	// search for "S" or "|"
	if y == 0 {
		for i := 0; i < len(input[0]); i++ {
			if input[y][i] == "S" {
				drawBeam(i, y+1, input)
				return calculateBeam(i, y+1, input)
			}
		}
	}

	if input[y][x] == "^" {
		add := 0
		leftPos := fmt.Sprintf("%d,%d", y, x-1)
		rightPos := fmt.Sprintf("%d,%d", y, x+1)

		if _, ok := outputPositions[leftPos]; !ok {
			add++
			outputPositions[leftPos] = true
		}
		if _, ok := outputPositions[rightPos]; !ok {
			add++
			outputPositions[rightPos] = true
		}

		// drawBeam(x-1, y, input)
		// drawBeam(x+1, y, input)
		// uniqueCarat[fmt.Sprintf("%d,%d", y, x)] = true
		// fmt.Println("split at x: ", x, "y:", y)
		memo[fmt.Sprintf("%d,%d", y, x-1)] = calculateBeam(x-1, y, input)
		memo[fmt.Sprintf("%d,%d", y, x+1)] = calculateBeam(x+1, y, input)
		return memo[fmt.Sprintf("%d,%d", y, x-1)] + memo[fmt.Sprintf("%d,%d", y, x+1)]
	}
	drawBeam(x, y+1, input)
	return calculateBeam(x, y+1, input)
}

func main() {
	input := readInput()
	numTimes := calculateBeam(0, 0, input)
	// printDiagram(input)
	fmt.Println(numTimes)
}
