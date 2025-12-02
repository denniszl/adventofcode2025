package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() []string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(dat), "\n")
}

// should have used python
// https://stackoverflow.com/questions/43018206/modulo-of-negative-integers-in-go
func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {
	input := readInput()
	start := 50
	zeroCount := 0
	for _, i := range input {
		if len(strings.TrimSpace(string(i))) == 0 {
			break
		}
		switch string(i[0]) {
		case "R":
			strings := strings.Split(string(i), "R")
			movement, err := strconv.Atoi(strings[1])
			if err != nil {
				panic(err)
			}
			start += movement
			start = mod(start, 100)
			fmt.Println("input: ", i, "pointer: ", start)
		case "L":
			strings := strings.Split(string(i), "L")
			movement, err := strconv.Atoi(strings[1])
			if err != nil {
				panic(err)
			}
			start -= movement
			start = mod(start, 100)
			fmt.Println("input: ", i, "pointer: ", start)
		}

		if start == 0 {
			zeroCount++
		}
	}

	fmt.Println(zeroCount)
}
