package twosum

import (
	"fmt"
	"testing"
)

type testcase struct {
	nums   []int
	target int
	wanted []int
}

func TestTwosum(t *testing.T) {
	testcases := []testcase{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
		{
			[]int{9, 5, 23, 45, 3, -2, -5},
			0,
			[]int{1, 6},
		},
		{
			[]int{9, 5, 23, 3, 3, -2, -5},
			6,
			[]int{3, 4},
		},
	}
	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")
	for _, tc := range testcases {
		act := twoSum(tc.nums, tc.target)
		if act[0] != tc.wanted[0] || act[1] != tc.wanted[1] {
			t.Errorf("[Input]: %v, [Output]: %v, [Wanted]: %v\n", tc.nums, act, tc.wanted)
		}
	}
}
