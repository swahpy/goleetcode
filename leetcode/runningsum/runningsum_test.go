package runningsum

import (
	"fmt"
	"testing"
)

type testcase struct {
	nums []int
	res  []int
}

func TestRunningsum(t *testing.T) {
	testcases := []testcase{
		{
			[]int{1, 2, 3, 4},
			[]int{1, 3, 6, 10},
		},
		{
			[]int{1, 1, 1, 1, 1},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{3, 1, 2, 10, 1},
			[]int{3, 4, 6, 16, 17},
		},
	}
	fmt.Printf("------------------------Leetcode Problem 1480------------------------\n")
	for _, v := range testcases {
		r := runningSum(v.nums)
		for i, vi := range r {
			if vi != v.res[i] {
				t.Errorf("[Input]: %v, [Output]: %v, [Wanted]: %v\n", v.nums, r, v.res)
			}
		}
	}
}
