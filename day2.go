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

// BoxProximity calculate how close two boxes are
func BoxProximity(boxA string, boxB string) (int, string) {
	aRunes := []rune(boxA)
	bRunes := []rune(boxB)

	same := make([]rune, 0)
	var diff int
	for i := 0; i < len(aRunes); i++ {
		if aRunes[i] == bRunes[i] {
			same = append(same, aRunes[i])
		} else {
			diff++
		}
	}
	return diff, string(same)
}

// BoxSetProximity proximity of all the boxIDs
func BoxSetProximity(boxIDs ...string) string {
	var answer string

	length := len(boxIDs)
	var cycles int
	log.Printf("length: %d\n", length)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			cycles++
			diff, same := BoxProximity(boxIDs[i], boxIDs[j])
			if diff == 1 {
				log.Printf("found after %d comparisons\n", cycles)
				answer = same
				return answer
			}
		}
	}
	return answer
}
