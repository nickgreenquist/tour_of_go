package main

import (
	"fmt"
	"time"
)

var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- sec
}

func main() {
	// c = make(chan int)
	// go ready("Tea", 2)
	// go ready("Coffee", 1)
	// fmt.Println("I'm waiting")
	
	// <- c
	// <- c
	// fmt.Println("I'm done")

	m := map[string]string {"key2": "hello"}
	fmt.Println(m["key3"])
}