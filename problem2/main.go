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
	return strings.Split(string(dat), ",")
}

func atoi(in string) int64 {
	v, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("errant string: %s", in))
	}
	return v
}

func isRepeatedString(in int64) bool {
	// heuristic for speed, condition is not fulfillable if not even length
	windowSize := len(strconv.FormatInt(in, 10))
	if windowSize&1 == 0 {
		// fmt.Println("even size: ", windowSize, "input: ", in)
		s := strconv.FormatInt(in, 10)
		// fmt.Println("1st half: ", s[0:windowSize/2], "second half: ", s[windowSize/2:])
		return s[0:windowSize/2] == s[windowSize/2:]
	}

	return false
}

func findRepeats(l, h int64) []int64 {
	out := []int64{}
	for i := l; i <= h; i++ {
		if isRepeatedString(i) {
			out = append(out, i)
		}
	}
	return out
}

func main() {
	inputs := readInput()
	for i := range inputs {
		inputs[i] = strings.TrimSpace(inputs[i])
	}
	sum := int64(0)
	for _, input := range inputs {
		s := strings.Split(input, "-")
		lower, upper := atoi(s[0]), atoi(s[1])
		repeats := findRepeats(lower, upper)
		// if len(repeats) > 0 {
		// 	fmt.Println("input: ", input, "repeats", repeats)
		// }
		for _, r := range repeats {
			sum += r
		}
	}

	fmt.Println("sum: ", sum)
}
