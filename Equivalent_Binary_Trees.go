package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walk func(t *tree.Tree, ch chan int)
	walk = func(t *tree.Tree, ch chan int) {
		if t.Left != nil {
			walk(t.Left, ch)
		}
		ch <- t.Value
		if t.Right != nil {
			walk(t.Right, ch)
		}
	}
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v := range ch1 {
		if v != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(5), tree.New(8)))
	fmt.Println(Same(tree.New(5), tree.New(5)))
}
