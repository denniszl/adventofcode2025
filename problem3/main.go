package main

import (
	"fmt"
	"math"
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

func atoi(in string) int64 {
	v, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("errant string: %s", in))
	}
	return v
}

func max(a, b int64) int64 {
	return int64(math.Max(float64(a), float64(b)))
}

func getIntFromString(in []string) int64 {
	return atoi(strings.Join(in, ""))
}

func getMaxJoltage(in []string, currentQueue []string, idx int) int64 {
	// fmt.Println("currentQueue: ", currentQueue, "idx: ", idx)
	if idx == len(in) && len(currentQueue) != 2 {
		return 0
	}

	if len(currentQueue) >= 2 {
		return getIntFromString(currentQueue)
	}

	return max(getMaxJoltage(in, append(currentQueue, in[idx]), idx+1), getMaxJoltage(in, currentQueue, idx+1))
}

func main() {
	input := readInput()
	maxJoltage := int64(0)
	v := [][]string{}
	for _, i := range input {
		if strings.TrimSpace(i) == "" {
			continue
		}
		vv := []string{}
		for _, r := range i {
			vv = append(vv, string(r))
		}
		v = append(v, vv)
	}
	for _, j := range v {
		localJ := getMaxJoltage(j, []string{}, 0)
		maxJoltage += localJ
		fmt.Println("input: ", j, "max: ", localJ)
	}
	fmt.Println("max joltage: ", maxJoltage)
}
