package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func highestDigit(line string) (byte, int) {
	max := byte('0' - 1)
	index := -1
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' && line[i] > max {
			max = line[i]
			index = i
		}
	}
	return max, index
}

func findHighestJoltage2(line string) int {
	length := len(line) - 1
	digit, index := highestDigit(line[:length])
	digit2, index := highestDigit(line[index+1:])
	joltage, err := strconv.Atoi(string(digit) + string(digit2))
	errHandler(err)
	return joltage
}

func findHighestJoltagen(line string, n int) string {
	length := len(line) - n + 1
	digit, index := highestDigit(line[:length])
	if n == 1 {
		return string(digit)
	}

	return string(digit) + findHighestJoltagen(line[index+1:], n-1)
}

func readLinesFromFile(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total1 := 0
	total2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		total1 += findHighestJoltage2(line)
		joltage, err := strconv.Atoi(findHighestJoltagen(line, 12))
		errHandler(err)
		total2 += joltage
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return total1, total2
}

func advent3() {
	totalJoltage1, totalJoltage2 := readLinesFromFile("input3.txt")
	fmt.Printf("ADVENT DAY 3\nThe total Joltage for part 1 is --> %d and part 2 is --> %d\n", totalJoltage1, totalJoltage2)
}
