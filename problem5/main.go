package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Interval struct {
	start int64
	end   int64
}

func (i Interval) String() string {
	return fmt.Sprintf("{low: %d, high: %d}\n", i.start, i.end)
}

/*
{low: 400509640504551, high: 400956320724574}
{low: 400509640504551, high: 400956320724574}
*/
func (i Interval) IsSubset(check Interval) bool {
	// fmt.Println("old:", i, "check:", check)
	// if i.start == 400509640504551 {
	// 	fmt.Println("i.start:", i.start, "check.start:", check.start)
	// }
	return i.start <= check.start && check.end <= i.end
}

func readInput() ([]Interval, map[string]bool) {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n")
	intervals := []Interval{}
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
			split1 := strings.Split(line, "-")
			intervals = append(intervals, Interval{
				start: atoi(split1[0]),
				end:   atoi(split1[1]),
			})
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

func between(a, b, comparison int64) bool {
	return a <= comparison && comparison <= b
}

func main() {
	intervals, _ := readInput()
	seenIntervals := []Interval{}
	slices.SortFunc(
		intervals,
		func(a, b Interval) int {
			if a.start < b.start {
				return -1
			}
			if a.start > b.start {
				return 1
			}
			// Tiebreaker: If starts are equal (c == 0), sort by end
			if a.end < b.end {
				return -1
			}
			if a.end > b.end {
				return 1
			}
			return 0 // items are fully equal
		})
	fmt.Println(intervals)
	numFresh := int64(0)
	l, h := intervals[0].start, intervals[0].end
	// doFinalCalc := true
	for i, interval := range intervals {
		if i == 0 {
			continue
		}
		if interval.start > h {
			if i+1 == len(intervals) {
				// doFinalCalc = false
			}
			seen := false
			for _, interval := range seenIntervals {
				if interval.IsSubset(Interval{
					start: l,
					end:   h,
				}) == true {
					fmt.Println("old interval")
					seen = true
				}
			}
			if !seen {
				numFresh += h - l + 1
			}
			seenIntervals = append(seenIntervals, Interval{
				start: l,
				end:   h,
			})
			// fmt.Println("low", l, "high", h, "current interval start", interval.start, "current interval end", interval.end, "numFresh:", numFresh)
			l, h = interval.start, interval.end
			continue
		}
		if interval.end > h {
			// fmt.Println("moving end window, old one:", h, "new one:", interval.end, "start: ", l, "pointer start: ", interval.start, "distance increased by: ", interval.end-h)
			h = interval.end
		}
	}

	// if doFinalCalc {
	// 	fmt.Println("doing final calc")
	numFresh += h - l + 1
	// }

	fmt.Println(numFresh)
}
