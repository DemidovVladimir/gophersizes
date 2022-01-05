package main

import (
	"fmt"
	"sync"
)

type test struct{}

func sq(l []int) chan int {
	ch := make(chan int)

	go func() {
		for _, v := range l {
			ch <- v * v
		}
		close(ch)
	}()

	return ch
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	nums := []int{5, 2, 3, 4, 234, 14243, 346436, 346, 32, 1}

	c1 := sq(nums)
	c2 := sq(nums)

	for v := range merge(&wg, c1, c2) {
		fmt.Println(v)
	}

}

func merge(wg *sync.WaitGroup, ch ...<-chan int) <-chan int {
	o := make(chan int)
	for _, v := range ch {
		go func(c <-chan int) {
			for n := range c {
				o <- n
			}
			wg.Done()
		}(v)
	}

	go func() {
		wg.Wait()
		close(o)
	}()

	return o
}
