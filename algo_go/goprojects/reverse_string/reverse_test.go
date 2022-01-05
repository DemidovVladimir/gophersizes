package reversestring

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		test string
		want string
	}{
		{"test", "tset"},
		{"ca t", "t ac"},
		{"jordan, mike", "ekim ,nadroj"},
		{"", ""},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("(%v,%v)", tc.test, tc.want), func(t *testing.T) {
			got := ReverseString(tc.test)
			if got != tc.want {
				t.Fatalf("NumInList() = %v; want %v", got, tc.want)
			}
		})
	}

}
