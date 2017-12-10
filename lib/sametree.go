package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"time"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		x, y := <-ch1, <-ch2
		if x != y {
			fmt.Printf("x=%d, y=%d\n", x, y)
			return false
		}
	}
	return true
}

func main() {
	t := Same(tree.New(1), tree.New(1))
	fmt.Printf("first two trees are = %v\n", t)
	t = Same(tree.New(1), tree.New(2))
	fmt.Printf("NExt two trees are = %v\n", t)
}
