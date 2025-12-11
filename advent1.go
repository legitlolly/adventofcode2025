package main

import (
    "bufio"
    "fmt"
    "os"
	"log"
	"strconv"
)


func commandBuilder(s string) int {
    dir := s[0]
    n, err := strconv.Atoi(s[1:])
    if err != nil {
        return 0 
    }

    if dir == 'L' {
        return -(n)
    }
    return n
}

func solve() (int, int) {
    file, err := os.Open("input1.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	pointer := 50
	passcode := 0
	count2 := 0
    for scanner.Scan() {
		move := commandBuilder(scanner.Text())
		pointer = pointer + move
		if pointer > 99 {
			count2 += pointer / 100
		}
		if pointer < 1 {
			count2 += -move / 100
		}
		pointer = ((pointer % 100) + 100) % 100
		if pointer == 0 {
			passcode++
			count2++
		} 
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return passcode, count2

}


func advent1() {
	passcode, count2 := solve()
	fmt.Printf("ADVENT DAY 1\nPart 1 --> %d Part 2 --> %d\n", passcode, count2)
}