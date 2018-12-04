package main

import (
	"log"
	"strconv"
)

// ComputeFrequencies variadic sum function
func ComputeFrequencies(deltas ...int) []int {
	frequencies := make([]int, len(deltas)+1)

	for i, delta := range deltas {
		//		log.Printf("i: %d\n", i)
		newFrequency := frequencies[i] + delta
		frequencies[i+1] = newFrequency
	}

	return frequencies
}

func scanDay1File() []int {
	lines, err := FileInput(1)
	if err != nil {
		log.Fatalf("failed to get data, %s", err)
	}

	var numbers []int

	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("could not convert %s, %s\n", line, err)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func day1() int {
	numbers := scanDay1File()

	frequencies := ComputeFrequencies(numbers...)

	// Last item of slice is the final frequency
	return frequencies[len(frequencies)-1]
}

func day1pt2() int {
	numbers := scanDay1File()

	dup := FindFirstDuplicate(numbers...)

	return dup
}

// FindFirstDuplicate find first dupe
func FindFirstDuplicate(deltas ...int) int {
	// The first frequency is always 0, so make room for it
	freqs := make([]int, len(deltas)+1)
	freqs[0] = 0

	// This will need to expand as needed. Don't bother trying guess what the size will be
	seen := make(map[int]bool)

	dIdx := 0
	curFreq := freqs[0]
	var newFreq int
	var dup int
	var foundDup bool

	// Loop forever if needed
	for i := 0; foundDup == false; i++ {
		newFreq = curFreq + deltas[dIdx]

		_, ok := seen[curFreq]

		if ok {
			dup = curFreq
			foundDup = true
		} else {
			seen[curFreq] = true
		}

		curFreq = newFreq
		dIdx++
		if dIdx == len(deltas) {
			dIdx = 0 // Set the deltas index back to the beginning
		}
	}

	//	log.Printf("foundDup: %+v\n", foundDup)
	//log.Printf("dup: %+v\n", dup)
	//	log.Printf("freqs: %+v\n", freqs)
	//	log.Printf("seen: %+v\n", seen)

	return dup
}

/*
   +1, -1 first reaches 0 twice.
0, +1, -1
0,  1,  0

   +3, +3, +4, -2, -4 first reaches 10 twice.
0, +3, +3, +4, -2, -4, +3, +3
0,  3,  6,  10, 8,  4,  7, 10

   -6, +3, +8, +5, -6 first reaches 5 twice.
0, -6, +3, +8, +5,
0, -6, -3,  5,  5

   +7, +7, -2, -7, -4 first reaches 14 twice.
0, +7, +7, -2, -7, -4, +7, +7, -2, -7, -4, +7, +7, -2
0,  7, 14, 12,  5,  1,  8, 15, 13,  6,  2,  9, 16, 14
*/
