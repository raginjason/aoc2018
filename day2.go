package main

import (
	"log"
)

// BoxLetterCount checksum helper
func BoxLetterCount(boxID string) (int, int) {
	seen := make(map[rune]int)
	for _, char := range boxID {
		seen[char] = seen[char] + 1
	}

	var twice int
	var thrice int
	for _, cnt := range seen {
		if cnt == 2 {
			twice = 1
		}
		if cnt == 3 {
			thrice = 1
		}
	}

	return twice, thrice
}

// BoxSetCount checksum all the boxIDs
func BoxSetCount(boxIDs ...string) int {
	var twice int
	var thrice int
	for _, boxID := range boxIDs {
		a, b := BoxLetterCount(boxID)
		twice = twice + a
		thrice = thrice + b
	}
	return twice * thrice
}

func scanDay2File() []string {
	lines, err := FileInput(2)
	if err != nil {
		log.Fatalf("failed to get data, %s", err)
	}

	return lines
}

func day2() int {
	lines := scanDay2File()

	checksum := BoxSetCount(lines...)

	return checksum
}
