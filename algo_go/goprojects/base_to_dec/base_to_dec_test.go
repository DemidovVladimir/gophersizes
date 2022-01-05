package basetodec

import "testing"

func TestBaseToDec(t *testing.T) {
	testCases := []struct {
		input string
		base  int
		want  int
	}{
		{"1110", 2, 14},
		{"0001", 2, 1},
		{"FEE", 16, 4078},
	}

	for _, tc := range testCases {
		t.Run("Base case to decimal", func(t *testing.T) {
			got := BaseToDec(tc.input, tc.base)
			if got != tc.want {
				t.Fatalf("NumInList() = %v; want %v", got, tc.want)
			}
		})
	}
}
