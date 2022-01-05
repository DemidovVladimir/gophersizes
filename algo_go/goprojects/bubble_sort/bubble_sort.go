package bubblesort

func BubbleSort(li []int) []int {
	for i := 0; i < len(li); i++ {
		for j := 0; j < len(li); j++ {
			if li[i] < li[j] {
				li[i], li[j] = li[j], li[i]
			}
		}
	}
	return li
}
