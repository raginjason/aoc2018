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

func TestBoxProximity(t *testing.T) {
	testCases := []struct {
		boxA     string
		boxB     string
		wantDiff int
		wantSame string
	}{
		{"abcde", "axcyd", 2, "ace"},
		{"fghij", "fguij", 1, "fgij"},
	}

	for _, tc := range testCases {
		t.Run(tc.boxA, func(t *testing.T) {
			gotDiff, gotSame := BoxProximity(tc.boxA, tc.boxB)
			if gotDiff != tc.wantDiff {
				t.Errorf("got %d; want %d", gotDiff, tc.wantDiff)
			}
			if gotSame != tc.wantSame {
				t.Errorf("got %s; want %s", gotSame, tc.wantSame)
			}
		})
	}
}
