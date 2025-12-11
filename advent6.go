package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Transpose(a [][]string) [][]string {
	if len(a) == 0 {
		return [][]string{}
	}

	rows := len(a)
	cols := len(a[0])

	result := make([][]string, cols)
	for i := range cols {
		result[i] = make([]string, rows)
	}

	for r := range rows {
		for c := range cols {
			result[c][r] = a[r][c]
		}
	}

	return result
}

func LoadMathGrid(path string) [][]string {
	file, err := os.Open(path)
	errHandler(err)
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		grid = append(grid, values)
	}

	return Transpose(grid)
}

func advent6() {
	grid := LoadMathGrid("input6.txt")
	//total := CalculateGridMath(grid)
	fmt.Print(grid)
}
