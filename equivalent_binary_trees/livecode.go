package main

import (
	"fmt"
	"time"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	walk(t, ch)
}

func walk(t *tree.Tree, ch chan<- int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	time.Sleep(3 * time.Second)
	ch <- t.Value
	walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	// for v1 := range c1 {
	// 	v2 := <-c2
	// 	fmt.Printf("c1 value is: %d, c2 value is: %d \n", v1, v2)
	// 	if v1 != v2 {
	// 		return false
	// 	}
	// }

	// case1:
	// 1, 2, 3, 4, 5, 0, 0, 0
	// 1, 2, 3, 4, 5

	// case2:
	// 1, 2, 3, 4, 5
	// 1, 2, 3, 4, 5, 6, 7
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		fmt.Printf("c1 value is: %d, c2 value is: %d \n", v1, v2)
		if ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}

		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	// example one
	fmt.Println(Same(tree.New(5), tree.New(8)))
	// example two
	start := time.Now()
	fmt.Println(Same(tree.New(5), tree.New(5)))
	fmt.Println(time.Now().Sub(start))
}
