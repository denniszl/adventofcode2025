package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() ([]string, []string) {
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
	return intervals, ingredients
}

func atoi(in string) int64 {
	v, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("errant string: %s", in))
	}
	return v
}

func main() {
	intervals, ingredients := readInput()
	freshCount := 0
	intervalMap := map[int64]bool{}
	for _, interval := range intervals {
		splitted := strings.Split(interval, "-")
		for i := atoi(splitted[0]); i <= atoi(splitted[1]); i++ {
			intervalMap[i] = true
		}
	}

	for _, i := range ingredients {
		if _, ok := intervalMap[atoi(i)]; ok {
			freshCount++
			fmt.Println(i)
		}
	}

	fmt.Println(freshCount)
}
