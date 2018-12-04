package main

import (
	"fmt"
	"strconv"
	"strings"
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
			frequencies := ComputeFrequencies(tc.x, tc.y, tc.z)
			got := frequencies[len(frequencies)-1]
			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}

func TestFindFirstDuplicate(t *testing.T) {
	testCases := []struct {
		list []int
		want int
	}{
		{[]int{+1, -1}, 0},
		{[]int{+3, +3, +4, -2, -4}, 10},
		{[]int{-6, +3, +8, +5, -6}, 5},
		{[]int{+7, +7, -2, -7, -4}, 14},
	}

	for _, tc := range testCases {
		var strs []string
		for _, num := range tc.list {
			s := strconv.Itoa(num)
			strs = append(strs, s)
		}
		name := strings.Join(strs, "_")
		t.Run(name, func(t *testing.T) {
			got := FindFirstDuplicate(tc.list...)
			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}
