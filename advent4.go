package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var adjacent8 = [][2]int{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func LoadGrid(path string) [][]uint8 {
	file, err := os.Open(path)
	errHandler(err)
	defer file.Close()

	var grid [][]uint8
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]uint8, len(line))

		for i, c := range line {
			switch c {
			case '.':
				row[i] = 0
			case '@':
				row[i] = 1
			default:
				log.Fatal("Unexpected Character")
			}
		}
		grid = append(grid, row)
	}
	return grid
}

func FindMoveableRolls(grid [][]uint8, height int, width int) int {
	collectibleCount := 0
	for y := range height {
		for x := range width {
			if grid[y][x] == 0 {
				continue
			}
			totalNeighbourPapers := 0
			for _, n := range adjacent8 {
				nx := x + n[0]
				ny := y + n[1]

				if nx < 0 || nx >= width || ny < 0 || ny >= height {
					continue
				}

				totalNeighbourPapers += int(grid[ny][nx])
			}
			if totalNeighbourPapers < 4 {
				collectibleCount++
			}
		}
	}
	return collectibleCount
}

func MaxFindMoveableRolls(grid [][]uint8, height int, width int) int {
	collectibleCount := 0
	for y := range height {
		for x := range width {
			if grid[y][x] == 0 {
				continue
			}
			totalNeighbourPapers := 0
			for _, n := range adjacent8 {
				nx := x + n[0]
				ny := y + n[1]

				if nx < 0 || nx >= width || ny < 0 || ny >= height {
					continue
				}

				totalNeighbourPapers += int(grid[ny][nx])
			}
			if totalNeighbourPapers < 4 {
				collectibleCount++
				grid[y][x] = 0
			}
		}
	}
	if collectibleCount == 0 {
		return 0
	}
	return collectibleCount + MaxFindMoveableRolls(grid, height, width)
}

func advent4() {
	grid := LoadGrid("input4.txt")
	height := len(grid)
	width := len(grid[0])

	moveableRolls := FindMoveableRolls(grid, height, width)
	maxMoveableRolls := MaxFindMoveableRolls(grid, height, width)

	fmt.Printf("ADVENT DAY 4\nThe count for moveable papers in part 1 --> %d and part 2 --> %d\n", moveableRolls, maxMoveableRolls)

}
