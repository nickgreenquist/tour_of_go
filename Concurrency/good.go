package main

import (
	"fmt"
)

func print(i int, c chan int) {
	fmt.Println(i)

	// this signals to the channel this routine is done
	c <- 1
}

func main() {
	num_calls := 20
	c := make(chan int, num_calls)

	for i := 0; i < num_calls; i++ {
		go print(i, c)
	}

	// this loop won't terminate until 20 ints have been popped out
	for i := 0; i < num_calls; i++ {
		<- c
	}
}