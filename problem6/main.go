package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func atoi(in string) (int64, error) {
	return strconv.ParseInt(in, 10, 64)
}

func readInput() (numbers [][]int64, operation []string) {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	splitted := strings.Split(string(dat), "\n")
	for _, v := range splitted {
		if strings.TrimSpace(v) == "" {
			break
		}
		splitted1 := strings.Fields(v)
		_, err := atoi(splitted1[0])
		if err != nil {
			operation = splitted1
			break
		}
		for j, sv := range splitted1 {
			vv, err := atoi(sv)
			if err != nil {
				panic(err)
			}
			if j == len(numbers) {
				numbers = append(numbers, []int64{})
				numbers[j] = []int64{}
			}
			numbers[j] = append(numbers[j], vv)
		}
	}

	return
}

func main() {
	numbers, operations := readInput()
	if len(numbers) != len(operations) {
		// fmt.Println("len numbers:", len(numbers), "len operations:", len(operations))
		panic("not equal length")
	}
	runningSum := int64(0)
	for i, arr := range numbers {
		localSum := int64(0)
		for j, num := range arr {
			if j == 0 {
				localSum = num
				continue
			}
			switch operations[i] {
			case "*":
				localSum = localSum * num
			case "+":
				localSum = localSum + num
			}
		}
		runningSum += int64(localSum)
	}
	fmt.Println(runningSum)
}
