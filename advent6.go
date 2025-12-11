package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func Multiply(x, y int) int {
	return x * y
}

func Add(x, y int) int {
	return x + y
}

func CalculateGridMath(grid [][]string) int {
	total := 0
	var loopTotal int
	var f func(x, y int) int
	for _, equation := range grid {
		length := len(equation) - 1
		if equation[length] == "*" {
			f = Multiply
			loopTotal = 1
		} else if equation[length] == "+" {
			f = Add
			loopTotal = 0
		} else {
			log.Fatal("WOAH WHAT IS GOING ON HERE ---> UNEXPECTED MATH OPERATION")
		}
		for _, value := range equation[:length] {
			value, err := strconv.Atoi(value)
			errHandler(err)
			loopTotal = f(value, loopTotal)
		}
		total += loopTotal
	}
	return total
}

func LoadWeirdCephalopodReadingStyle(grid [][]string) [][]string {
	fmt.Print(grid)
	return grid
}

func advent6() {
	grid := LoadMathGrid("input6.txt")
	total := CalculateGridMath(grid)
	grid = LoadWeirdCephalopodReadingStyle(grid)
	fmt.Print(total)
}
