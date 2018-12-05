package main

import (
	"log"
	"testing"
)

func TestNewClaim(t *testing.T) {
	testCases := []struct {
		input string
		want  Claim
	}{
		{"#123 @ 3,2: 5x4", Claim{claimID: 123, leftOffset: 3, topOffset: 2, width: 5, height: 4}},
		{"#1 @ 1,3: 4x4", Claim{claimID: 1, leftOffset: 1, topOffset: 3, width: 4, height: 4}},
		{"#2 @ 3,1: 4x4", Claim{claimID: 2, leftOffset: 3, topOffset: 1, width: 4, height: 4}},
		{"#3 @ 5,5: 2x2", Claim{claimID: 3, leftOffset: 5, topOffset: 5, width: 2, height: 2}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			c := NewClaim(tc.input)
			if c.claimID != tc.want.claimID {
				t.Errorf("got %d; want %d", c.claimID, tc.want.claimID)
			}
			if c.leftOffset != tc.want.leftOffset {
				t.Errorf("got %d; want %d", c.leftOffset, tc.want.leftOffset)
			}
			if c.topOffset != tc.want.topOffset {
				t.Errorf("got %d; want %d", c.topOffset, tc.want.topOffset)
			}
			if c.width != tc.want.width {
				t.Errorf("got %d; want %d", c.width, tc.want.width)
			}
			if c.height != tc.want.height {
				t.Errorf("got %d; want %d", c.height, tc.want.height)
			}
		})
	}
}

func TestGrid(t *testing.T) {
	testCases := []struct {
		input string
		want  Claim
	}{
		{"#1 @ 1,3: 4x4", Claim{claimID: 1, leftOffset: 1, topOffset: 3, width: 4, height: 4}},
		{"#2 @ 3,1: 4x4", Claim{claimID: 2, leftOffset: 3, topOffset: 1, width: 4, height: 4}},
		{"#3 @ 5,5: 2x2", Claim{claimID: 3, leftOffset: 5, topOffset: 5, width: 2, height: 2}},
		{"#123 @ 3,2: 5x4", Claim{claimID: 123, leftOffset: 3, topOffset: 2, width: 5, height: 4}},
	}

	g := make([][]int, 0, 0)
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			c := NewClaim(tc.input)
			g = c.grid(g)
			log.Printf("contents of claim %d (top: %d, height: %d, left: %d, width: %d)\n", c.claimID, c.topOffset, c.height, c.leftOffset, c.width)
			for i := range g {
				log.Printf("%v\n", g[i])
			}
			log.Printf("%#v\n", g)
		})
	}
}

func TestCountIntersection(t *testing.T) {
	testCases := []struct {
		input [][]int
		want  int
	}{
		{[][]int{[]int(nil), []int{0, 0, 0, 1, 1, 1, 1}, []int{0, 0, 0, 1, 1, 1, 1}, []int{0, 1, 1, 2, 2, 1, 1}, []int{0, 1, 1, 2, 2, 1, 1}, []int{0, 1, 1, 1, 1, 1, 1}, []int{0, 1, 1, 1, 1, 1, 1}}, 4},
	}

	for _, tc := range testCases {
		t.Run("A", func(t *testing.T) {
			got := CountIntersection(tc.input)
			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}
