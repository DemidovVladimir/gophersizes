package fizbuz

import "fmt"

func FizBuz(i int) string {
	if i%3 == 0 {
		return "fiz"
	} else if i%5 == 0 {
		return "buz"
	} else {
		return fmt.Sprint(i)
	}
}
