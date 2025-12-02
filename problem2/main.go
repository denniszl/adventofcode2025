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

func hasRepeatedString(in int64) bool {
	windowSize := len(strconv.FormatInt(in, 10))
	stringInt := strconv.FormatInt(in, 10)
	for i := 1; i < windowSize/2+1; i++ {
		fmt.Println("checking substr: ", stringInt[0:i], "against: ", stringInt)
		if numRepeats(stringInt[0:i], stringInt) >= 2 {
			return true
		}
	}

	return false
}
func numRepeats(substr, str string) int {
	if substr == "" || str == "" {
		return 0
	}
	// fmt.Println("substr: ", substr, "string: ", str)
	if len(substr) > len(str) || (len(substr) == len(str) && substr != str) {
		// should be intmax tbh, but let's do FF numbers
		return -9999999
	}
	if substr == str {
		return 1
	}
	before, after, found := strings.Cut(str, substr)
	if before == "" && after == "" && found {
		return 1
	}
	if before != "" {
		return -9999999
	}
	if !found {
		return 0
	}
	return 1 + numRepeats(substr, after)
}

func findRepeats(l, h int64) []int64 {
	out := []int64{}
	for i := l; i <= h; i++ {
		if hasRepeatedString(i) {
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
		fmt.Println("lower: ", lower, "upper: ", upper)
		repeats := findRepeats(lower, upper)
		if len(repeats) > 0 {
			fmt.Println("input: ", input, "repeats", repeats)
		}
		for _, r := range repeats {
			sum += r
		}
	}

	fmt.Println("sum: ", sum)
}
