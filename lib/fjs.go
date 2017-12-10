package main

import (
	"fmt"
)

func main() {
	var d DisjointSet
	d.init(10)
	d.WeightedUnion(1, 2)
	d.Print()
	fmt.Println("vim-go")
}
