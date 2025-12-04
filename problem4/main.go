package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() [][]string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(dat), "\n")
	out := [][]string{}
	for _, splitted := range split {
		if strings.TrimSpace(splitted) == "" {
			continue
		}
		str := []string{}
		for _, r := range splitted {
			str = append(str, string(r))
		}
		out = append(out, str)
	}

	return out
}

func prettyp(str [][]string) {
	a := []string{}
	for _, s := range str {
		a = append(a, strings.Join(s, ""))
	}
	fmt.Print(strings.Join(a, "\n"))
}

func prettypi(str [][]int) {
	a := []string{}
	for _, s := range str {
		for _, j := range s {
			a = append(a, strconv.FormatInt(int64(j), 10))
		}
	}
	fmt.Print(strings.Join(a, ""))
}

/*
	The forklifts can only access a roll of paper if there are fewer
	than four rolls of paper in the eight adjacent positions.
	If you can figure out which rolls of paper the
	forklifts can access, they'll spend less time
	looking and more time breaking down the wall to the cafeteria.
*/

// 111 i0[0][1][2]
// 111 i1[0][1][2]

var coordFilter = map[string][]int{
	"UL": []int{-1, -1},
	"U":  []int{-1, 0},
	"UR": []int{-1, 1},
	"L":  []int{0, -1},
	"R":  []int{0, 1},
	"DL": []int{1, -1},
	"D":  []int{1, 0},
	"DR": []int{1, 1},
}

func calculateNumberOfRollsForIndex(rolls [][]string, y, x int) int {
	if rolls[y][x] != "@" {
		return 9999
	}

	surroundedBy := 0
	maxY := len(rolls)
	maxX := len(rolls[0])
	for _, filter := range coordFilter {
		checkY, checkX := y+filter[0], x+filter[1]
		if checkY < 0 || checkX < 0 || checkY >= maxY || checkX >= maxX {
			continue
		}
		if rolls[checkY][checkX] == "@" {
			surroundedBy += 1
		}
	}
	// fmt.Println("surroundedBy: ", surroundedBy, "coordinate: x,y:", x, y)
	return surroundedBy
}

func main() {
	out := readInput()
	// fmt.Println(out)
	prettyp(out)
	fmt.Println("")
	countBelow4 := 0
	weights := [][]int{}
	for i := range out {
		ap := []int{}
		for j := range out[i] {
			weight := calculateNumberOfRollsForIndex(out, i, j)
			if weight < 4 {
				countBelow4++
			}
			ap = append(ap, weight)
		}
		weights = append(weights, ap)
	}

	fmt.Println(countBelow4)
}
