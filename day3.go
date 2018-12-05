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
