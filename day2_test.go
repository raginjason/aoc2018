package main

import (
	"testing"
)

func TestBoxLetterCount(t *testing.T) {
	testCases := []struct {
		boxID      string
		wantTwice  int
		wantThrice int
	}{
		{"abcdef", 0, 0},
		{"bababc", 1, 1},
		{"abbcde", 1, 0},
		{"abcccd", 0, 1},
		{"aabcdd", 1, 0},
		{"abcdee", 1, 0},
		{"ababab", 0, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.boxID, func(t *testing.T) {
			gotTwice, gotThrice := BoxLetterCount(tc.boxID)
			if gotTwice != tc.wantTwice {
				t.Errorf("got %d; want %d", gotTwice, tc.wantTwice)
			}
			if gotThrice != tc.wantThrice {
				t.Errorf("got %d; want %d", gotThrice, tc.wantThrice)
			}
		})
	}
}

func TestBoxSetCount(t *testing.T) {
	testCases := []struct {
		boxIDs []string
		want   int
	}{
		{[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, 12},
	}

	for _, tc := range testCases {
		t.Run(tc.boxIDs[0], func(t *testing.T) {
			got := BoxSetCount(tc.boxIDs...)
			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}
