package merge_search

import (
	"math/rand"
	"time"
)

func GenerateSlice() []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(10)
}

func MergeSort(items []int) []int {
	if len(items) == 0 {
		return []int{}
	}
	if len(items) == 1 {
		return items
	}

	median := len(items) / 2
	left := items[:median]
	right := items[median:]

	return Merge(MergeSort(left), MergeSort(right))
}

func Merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			result[i] = right[0]
			right = right[1:]
		} else if right[0] > left[0] {
			result[i] = left[0]
			left = left[1:]
		}
		i++
	}

	for l := 0; l < len(left); l++ {
		result[i] = left[l]
		i++
	}

	for l := 0; l < len(right); l++ {
		result[i] = right[l]
		i++
	}

	return
}
