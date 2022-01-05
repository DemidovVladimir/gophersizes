package bubblesort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 2, 1, 5, 4}, []int{1, 2, 3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run("Bubble sort algo...", func(t *testing.T) {
			got := BubbleSort(tc.input)
			for i, v := range got {
				if v != tc.want[i] {
					t.Fatalf("NumInList() = %v; want %v", got, tc.want)
				}
			}
		})
	}
}

func BenchmarkBubbleeSort(b *testing.B) {
	for _, size := range []int{
		100, 200, 400, 800, 1600,
	} {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				b.StopTimer()
				list := rand.Perm(size)
				b.StartTimer()
				BubbleSort(list)
			}
		})
	}
}
