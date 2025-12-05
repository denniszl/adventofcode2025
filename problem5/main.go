package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() ([]string, map[string]bool) {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n")
	intervals := []string{}
	ingredients := []string{}
	doIngredients := false
	for _, line := range lines {
		if strings.TrimSpace(line) == "" && doIngredients == false {
			doIngredients = true
			continue
		}
		if doIngredients == true {
			if strings.TrimSpace(line) == "" {
				break
			}
			ingredients = append(ingredients, line)
		} else {
			intervals = append(intervals, line)
		}
	}
	ingredientMap := map[string]bool{}
	for _, i := range ingredients {
		ingredientMap[i] = true
	}
	return intervals, ingredientMap
}

func atoi(in string) int64 {
	v, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("errant string: %s", in))
	}
	return v
}

func betweenInterval(low, high, cmp int64) bool {
	return low <= cmp && cmp <= high
}

func main() {
	intervals, _ := readInput()
	intervalMap := map[int64]bool{}
	for _, interval := range intervals {
		splitted := strings.Split(interval, "-")
		for i := atoi(splitted[0]); i <= atoi(splitted[1]); i++ {
			intervalMap[i] = true
		}
	}

	fmt.Println(len(intervalMap))
}
