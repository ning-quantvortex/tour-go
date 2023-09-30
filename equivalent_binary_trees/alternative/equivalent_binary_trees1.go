package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t.Left != nil {
			walker(t.Left)
		}
		ch <- t.Value
		if t.Right != nil {
			walker(t.Right)
		}
	}
	walker(t)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		n1, ok1 := <-ch1
		n2, ok2 := <-ch2
		fmt.Printf("ch1 value is: %d \n", n1)
		fmt.Printf("ch2 value is: %d \n", n2)
		if ok1 == false && ok2 == false {
			break
		}

		if ok1 != ok2 || n1 != n2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(5), tree.New(8)))
	fmt.Println(Same(tree.New(5), tree.New(5)))
}
