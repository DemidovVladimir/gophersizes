package reversestring

import "strings"

func ReverseString(s string) string {
	var nS strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		nS.WriteByte(s[i])
	}
	return nS.String()
}
