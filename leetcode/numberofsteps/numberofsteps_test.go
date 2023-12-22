package numberofsteps

import (
	"fmt"
	"testing"
)

type testcase struct {
	num  int
	want int
}

func TestNumberofsteps(t *testing.T) {
	testcases := []testcase{
		{14, 6},
		{8, 4},
		{123, 12},
		{0, 0},
	}
	fmt.Printf("------------------------Leetcode Problem 1342------------------------\n")
	for _, tc := range testcases {
		res := numberOfSteps(tc.num)
		if res != tc.want {
			t.Errorf("[Input]: %v, [Output]: %v, [Wanted]: %v\n", tc.num, res, tc.want)
		}
	}
}
