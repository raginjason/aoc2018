package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		x    int
		y    int
		z    int
		want int
	}{
		{+1, +1, +1, 3},
		{+1, +1, -2, 0},
		{-1, -2, -3, -6},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d+%d+%d", tc.x, tc.y, tc.z), func(t *testing.T) {
			got := Sum(tc.x, tc.y, tc.z)
			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}
