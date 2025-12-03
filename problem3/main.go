package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

type memoizationStruct struct {
	currentQueue string
	idx          int
}

var memo = map[memoizationStruct]int64{}

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

func getStringArrayFromInt(in int64) []string {
	s := strconv.FormatInt(in, 10)
	out := []string{}
	for _, r := range s {
		out = append(out, string(r))
	}
	return out
}

func checkMemo(currentQueue []string, idx int) (int64, bool) {
	// fmt.Println("memo hit: ", currentQueue, idx)
	v, ok := memo[memoizationStruct{
		currentQueue: strings.Join(currentQueue, ""),
		idx:          idx,
	}]
	return v, ok
}

func storeAndReturn(currentQueue []string, idx int, value int64) int64 {
	memo[memoizationStruct{
		currentQueue: strings.Join(currentQueue, ""),
		idx:          idx,
	}] = value
	return value
}

var currentMax = int64(0)

func getMaxJoltage(in []string, currentQueue []string, idx int) int64 {
	// fmt.Println("currentQueue: ", currentQueue, "idx: ", idx)
	v, ok := checkMemo(currentQueue, idx)
	if ok {
		return v
	}

	// abort early if it's clear this is a dud from any digit.
	if currentMax != 0 && len(currentQueue) > 0 {
		cms := getStringArrayFromInt(currentMax)
		j := 0
		for j = 0; j < len(currentQueue); j++ {
			if cms[j] != currentQueue[j] {
				break
			}
		}
		if j <= len(currentQueue)-1 {
			if cms[j] > currentQueue[j] {
				return storeAndReturn(currentQueue, idx, 0)
			}
		}
	}

	if len(currentQueue) >= 12 {
		if getIntFromString(currentQueue) > currentMax {
			currentMax = getIntFromString(currentQueue)
		}
		return storeAndReturn(currentQueue, idx, getIntFromString(currentQueue))
	}

	if len(currentQueue)+(len(in)-idx) < 12 {
		return storeAndReturn(currentQueue, idx, 0)
	}

	if idx == len(in) {
		return storeAndReturn(currentQueue, idx, 0)
	}

	return storeAndReturn(currentQueue, idx, max(getMaxJoltage(in, append(currentQueue, in[idx]), idx+1), getMaxJoltage(in, currentQueue, idx+1)))
}

func main() {
	defer timeTrack(time.Now(), "main")
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
		memo = map[memoizationStruct]int64{}
		currentMax = 0
		localJ := getMaxJoltage(j, []string{}, 0)
		maxJoltage += localJ
		fmt.Println("input: ", j, "max: ", localJ)
	}
	fmt.Println("max joltage: ", maxJoltage)
}
