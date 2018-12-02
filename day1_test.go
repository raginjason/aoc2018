package main

import "testing"

func TestSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		z int
		n int
	}{
		{+1, +1, +1, 3},
		{+1, +1, -2, 0},
		{-1, -2, -3, -6},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y, table.z)
		if total != table.n {
			t.Errorf("Sum of (%d+%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, table.z, total, table.n)
		}
	}
}
