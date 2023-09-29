package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

/* questions:
1. so normally sender close the chan?
2. should we always give the direction of the chan?
3.
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// defer or nor defer
	defer close(ch)
	var walk func(t *tree.Tree)
	// here is a closure for ch
	// we can use walk func(t *tree.Tree, ch) also
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v := range ch1 {
		v2 := <-ch2
		if v != v2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(5), tree.New(8)))
	fmt.Println(Same(tree.New(5), tree.New(5)))
}
