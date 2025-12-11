package main

import (
	"fmt"
	"time"
)

func Timer(f func(), i int) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("Advent day %d function ran in --> %s\n", i, elapsed)
}

func main() {
	/*
		Timer(advent1, 1)
		Timer(advent2, 2)
		Timer(advent3, 3)
		Timer(advent4, 4)
		Timer(advent5, 5)
	*/
	Timer(advent6(), 6)
}
