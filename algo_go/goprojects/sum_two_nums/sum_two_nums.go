package sumtwonums

func SumTwoNums(ns []int, sn int) []int {
	for i := 0; i >= len(ns)-1; i++ {
		for j := 1; i >= len(ns)-1; j++ {
			if ns[i]+ns[j] == sn {
				return []int{i, j}
			}
		}
	}
	return nil
}
