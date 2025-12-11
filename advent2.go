package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func errHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isMirrored(i int) int {
	str := strconv.Itoa(i)
	half := len(str) / 2
	if str[0:half] == str[half:] {
		return i
	}
	return 0
}

func isRepeated(i int) int {
	str := strconv.Itoa(i)
	n := len(str)
	doubled := (str + str)[1 : 2*n-1]
	if strings.Contains(doubled, str) {
		return i
	}
	return 0
}

func findMirrorNumbers(s string) (int, int) {
	split := strings.Split(s, "-")
	total2 := 0
	total1 := 0
	start, err := strconv.Atoi(split[0])
	errHandler(err)
	end, err := strconv.Atoi(split[1])
	errHandler(err)
	for i := start; i <= end; i++ {
		total1 += isMirrored(i)
		total2 += isRepeated(i)
	}
	return total1, total2
}

func advent2() {
	file, err := os.Open("input2.csv")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	total1 := 0
	total2 := 0
	for _, record := range records[0] {
		toadd1, toadd2 := findMirrorNumbers(record)
		total1 += toadd1
		total2 += toadd2
	}

	fmt.Printf("ADVENT DAY 2\nThe total count for part 1 is --> %d and for part 2 is --> %d\n", total1, total2)
}
