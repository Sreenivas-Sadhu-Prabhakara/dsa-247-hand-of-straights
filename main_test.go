package main

import "testing"

func TestIsNStraightHand(t *testing.T) {
	cases := []struct {
		hand      []int
		groupSize int
		want      bool
	}{
		{[]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3, true}, // 123, 234, 678
		{[]int{1, 2, 3, 4, 5}, 4, false},            // size doesn't divide
		{[]int{1, 2, 3}, 3, true},
		{[]int{1, 1, 2, 2, 3, 3}, 3, true}, // two parallel runs
		{[]int{1, 1, 2, 3}, 2, false},      // second 1 has no partner 2 left... 1-2,1-3? 3 not consecutive to 1: false
		{[]int{4}, 1, true},                // groups of one: always
		{[]int{}, 3, true},
		{[]int{1, 3}, 2, false}, // gap
	}
	for _, c := range cases {
		if got := IsNStraightHand(c.hand, c.groupSize); got != c.want {
			t.Errorf("IsNStraightHand(%v,%d) = %v, want %v", c.hand, c.groupSize, got, c.want)
		}
	}
}
