package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// var CLOSEA = false
var DATA = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int, signal <-chan struct{}) {
	for {
		select {
		case <-signal:
			close(out)
			return
		case out <- random(min, max):
		}
	}
}

func second(out chan<- int, in <-chan int, signal chan<- struct{}) {
	for x := range in {
		_, ok := DATA[x]
		if ok {
			signal <- struct{}{}
			fmt.Print(x, " ")
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	var sum int
	sum = 0
	for x2 := range in {
		sum = sum + x2
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)
}

func main() {
	var signal chan struct{}
	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters!")
		return
	}
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])
	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d\n", n1, n2)
		return
	}
	rand.Seed(time.Now().UnixNano())
	A := make(chan int)
	B := make(chan int)
	signal = make(chan struct{})
	go first(n1, n2, A, signal)
	go second(B, A, signal)
	third(B)
}
