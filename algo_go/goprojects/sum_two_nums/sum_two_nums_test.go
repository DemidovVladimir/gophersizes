package sumtwonums

import "testing"

func TestSumTwoNums(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
		want   []int
	}{
		{[]int{1, 2, 3, 4, 5}, 7, []int{1, 4}},
		{[]int{1, 2, 3, 4, 5}, 9, []int{3, 4}},
	}

	for _, tc := range testCases {
		t.Run("Base case to decimal", func(t *testing.T) {
			got := SumTwoNums(tc.input, tc.output)
			for i, v := range got {
				if v != tc.want[i] {
					t.Fatalf("NumInList() = %v; want %v", got, tc.want)
				}
			}
		})
	}
}
