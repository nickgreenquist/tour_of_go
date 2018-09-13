package main

import (
	"fmt"
)

func print(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i < 20; i++ {
		go print(i)
	}
}