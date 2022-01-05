package main

import (
	"fmt"
	"math/rand"
	"time"
)

func binarysearch(arr []int, s int) bool {
	if len(arr) == 0 {
		return false
	}

	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2

		if arr[median] < s {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != s {
		return false
	}

	return true
}

func main() {

	type readS struct {
		key  int
		resp chan int
	}

	type writeS struct {
		key   int
		value int
		resp  chan int
	}

	reads := make(chan readS)
	writes := make(chan writeS)

	go func() {
		var state = make(map[int]int)

		for {
			select {
			case r := <-reads:
				fmt.Println(r, " - - - reading data...")
				r.resp <- state[r.key]
			case w := <-writes:
				fmt.Println(w, " - - - wrigin data...")
				state[w.key] = w.value
				w.resp <- w.key
			}
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				read := readS{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				write := writeS{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					resp:  make(chan int),
				}
				writes <- write
				<-write.resp
				time.Sleep(2 * time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	// chani := make(chan int)

	// chani := go theads.Test()

	// a := <-chani

	// fmt.Println(a)
}

func handle(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
