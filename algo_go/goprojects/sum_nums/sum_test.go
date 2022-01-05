package sumnums

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		list []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{2, 3, -2}, 3},
		{[]int{0, 2, -3}, -1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("(%v,%v)", tc.list, tc.want), func(t *testing.T) {
			got := Sum(tc.list)
			if got != tc.want {
				t.Fatalf("NumInList() = %v; want %v", got, tc.want)
			}
		})
	}

}
