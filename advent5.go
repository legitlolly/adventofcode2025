package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func GetRanges(path string) ([]Range, []int) {
	file, err := os.Open(path)
	errHandler(err)
	defer file.Close()

	var ranges []Range
	var data []int
	rangeData := true
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			rangeData = false
			continue
		}
		if rangeData == true {
			parts := strings.Split(line, "-")
			start, err := strconv.Atoi(parts[0])
			errHandler(err)
			end, err := strconv.Atoi(parts[1])
			errHandler(err)
			ranges = append(ranges, Range{Start: start, End: end})
		} else {
			val, err := strconv.Atoi(line)
			errHandler(err)
			data = append(data, val)
		}
	}
	return ranges, data
}

func SortRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})
	return ranges
}

func CollapseRanges(ranges []Range) []Range {
	if len(ranges) <= 1 {
		return ranges
	}

	sortedRanges := SortRanges(ranges)

	collapsedRanges := []Range{sortedRanges[0]}

	for _, current := range sortedRanges[1:] {
		lastCollapsed := &collapsedRanges[len(collapsedRanges)-1]

		if current.Start <= lastCollapsed.End {
			if current.End > lastCollapsed.End {
				lastCollapsed.End = current.End
			}
		} else {
			collapsedRanges = append(collapsedRanges, current)
		}
	}
	return collapsedRanges
}

func FreshIngredients(ranges []Range, data []int) int {
	total := 0
	for x := range data {
		for y := range ranges {
			if data[x] >= ranges[y].Start && data[x] <= ranges[y].End {
				total++
				break
			}
		}
	}
	return total
}

func IdCount(ranges []Range) int {
	totalLen := 0
	for _, r := range ranges {
		totalLen += (r.End - r.Start) + 1
	}
	return totalLen
}

func advent5() {
	ranges, data := GetRanges("input5.txt")
	ranges = CollapseRanges(ranges)
	total := FreshIngredients(ranges, data)
	freshIds := IdCount(ranges)
	fmt.Printf("ADVENT DAY 5\nPart 1 --> %d Part 2 --> %d\n", total, freshIds)
}
