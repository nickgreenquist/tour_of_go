package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	_Walk(t, ch)
	close(ch)
}

func _Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		_Walk(t.Left, ch)
		ch <- t.Value
		_Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := range c1 {
		if i != <- c2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(2), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("Should return true:", Same(tree.New(1), tree.New(1)))
	fmt.Println("Should return false:", Same(tree.New(1), tree.New(2)))
}