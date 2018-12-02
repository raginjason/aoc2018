package main

import (
	"log"
	"strconv"
)

// Sum variadic sum function
func Sum(numbers ...int) int {
	var total int
	for _, number := range numbers {
		total = total + number
	}
	return total
}

func day1() int {
	lines, err := FileInput(1)
	if err != nil {
		log.Fatalf("failed to get data, %s", err)
	}
	var total int

	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("could not convert %s, %s\n", line, err)
		}
		log.Printf("current: %d, add: %d\n", total, num)
		total = Sum(total, num)
	}

	return total
}
