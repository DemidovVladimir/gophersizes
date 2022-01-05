package fizbuz

import (
	"fmt"
	"testing"
)

func TestFizBuzStrings(t *testing.T) {
	testcases := []struct {
		input int
		want  string
	}{
		{3, "fiz"},
		{5, "buz"},
		{2, "2"},
		{1, "1"},
	}
	for _, tc := range testcases {
		t.Run(fmt.Sprintf("(%v,%v)", tc.input, tc.want), func(t *testing.T) {
			got := FizBuz(tc.input)
			if got != tc.want {
				t.Fatalf("NumInList() = %v; want %v", got, tc.want)
			}
		})
	}
}
