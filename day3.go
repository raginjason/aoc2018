package main

import (
	"log"
	"strconv"
	"strings"
)

// Claim simple struct
type Claim struct {
	claimID    int
	leftOffset int
	topOffset  int
	width      int
	height     int
}

// NewClaim parse a new claim
func NewClaim(line string) *Claim {
	c := new(Claim)

	claimParts := strings.Split(line, " ")

	// ClaimID
	idStr := strings.TrimPrefix(claimParts[0], "#")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalf("could not convert %s, %s\n", idStr, err)
	}
	c.claimID = id

	offsets := strings.Split(strings.TrimSuffix(claimParts[2], ":"), ",")

	// Left
	leftOffset, err := strconv.Atoi(offsets[0])
	if err != nil {
		log.Fatalf("could not convert %s, %s\n", offsets[0], err)
	}
	c.leftOffset = leftOffset

	// Top
	topOffset, err := strconv.Atoi(offsets[1])
	if err != nil {
		log.Fatalf("could not convert %s, %s\n", offsets[1], err)
	}
	c.topOffset = topOffset

	sizes := strings.Split(claimParts[3], "x")

	// Width
	width, err := strconv.Atoi(sizes[0])
	if err != nil {
		log.Fatalf("could not convert %s, %s\n", sizes[0], err)
	}
	c.width = width

	// Height
	height, err := strconv.Atoi(sizes[1])
	if err != nil {
		log.Fatalf("could not convert %s, %s\n", sizes[1], err)
	}
	c.height = height

	return c
}

func (c Claim) grid(g [][]int) [][]int {

	if cap(g) < c.topOffset+c.height {
		newSize := c.topOffset + c.height
		newSlice := make([][]int, newSize, newSize)
		copy(newSlice, g)
		g = newSlice
	}

	for i := c.topOffset; i < c.topOffset+c.height; i++ {
		if cap(g[i]) < c.leftOffset+c.width {
			newSize := c.leftOffset + c.width
			newSlice := make([]int, newSize, newSize)
			copy(newSlice, g[i])
			g[i] = newSlice
		}
		for j := c.leftOffset; j < c.leftOffset+c.width; j++ {
			g[i][j]++
		}
	}
	return g
}

// CountIntersection squares
func CountIntersection(g [][]int) int {
	var count int
	for i := range g {
		for j := range g[i] {
			if g[i][j] > 1 {
				count++
			}
		}
	}
	return count
}

func scanDay3File() []string {
	lines, err := FileInput(3)
	if err != nil {
		log.Fatalf("failed to get data, %s", err)
	}

	return lines
}

func day3() int {
	lines := scanDay3File()

	grid := make([][]int, 0, 0)
	for _, line := range lines {
		claim := NewClaim(line)
		grid = claim.grid(grid)
	}

	intersection := CountIntersection(grid)

	return intersection
}
