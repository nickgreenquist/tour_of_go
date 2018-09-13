package main

import (
	"fmt"
	"sync"
)

func print(i int, wg *sync.WaitGroup) {
	// defer means run this line right before exiting the function
	// wg.Done() signals to WaitGroup that this routine is done
	defer wg.Done()
	
	fmt.Println(i)
}

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go print(i, wg)
	}

	wg.Wait()
}