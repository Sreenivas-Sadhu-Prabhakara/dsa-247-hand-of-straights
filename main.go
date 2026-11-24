package main

import (
	"fmt"
	"sort"
)

// IsNStraightHand reports whether the cards can be split entirely into
// groups of groupSize consecutive values. Greedy from the smallest card:
// it must start a group, so its whole run must be available.
func IsNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	count := map[int]int{}
	for _, v := range hand {
		count[v]++
	}
	values := make([]int, 0, len(count))
	for v := range count {
		values = append(values, v)
	}
	sort.Ints(values)
	for _, v := range values {
		c := count[v]
		if c == 0 {
			continue // already consumed as the tail of earlier groups
		}
		// The smallest remaining card starts c groups; take the runs whole.
		for k := v; k < v+groupSize; k++ {
			if count[k] < c {
				return false
			}
			count[k] -= c
		}
	}
	return true
}

func main() {
	fmt.Println(IsNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
	fmt.Println(IsNStraightHand([]int{1, 2, 3, 4, 5}, 4))
}
