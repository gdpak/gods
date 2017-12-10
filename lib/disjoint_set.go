package main

import (
	"fmt"
)

type DisjointSet struct {
	n      int
	rank   []int
	parent []int
}

func (d *DisjointSet) init(n int) {
	d.n = n
	d.rank = make([]int, n)
	d.parent = make([]int, n)
	for i := 0; i < n; i++ {
		d.parent[i] = i
	}
}

func (d *DisjointSet) Find(elem int) int {
	prev_parent := d.parent[elem]

	for d.parent[prev_parent] != prev_parent {
		prev_parent = d.parent[prev_parent]
	}
	d.parent[elem] = prev_parent
	return d.parent[elem]
}

func (d *DisjointSet) WeightedUnion(src int, dest int) int {
	var srcid int = d.Find(src)
	var destid int = d.Find(dest)
	var new_root int = 0

	switch {
	case d.rank[srcid] >= d.rank[destid]:
		d.parent[destid] = d.parent[srcid]
		new_root = d.parent[destid]
	default:
		d.parent[srcid] = d.parent[destid]
		new_root = d.parent[srcid]
	}
	if d.rank[srcid] == d.rank[destid] {
		d.rank[srcid] += 1
	}
	return new_root
}

func (d *DisjointSet) Print() {
	fmt.Printf("--- Disjoint Set- Parent list \n")
	for i := 0; i < d.n; i++ {
		fmt.Printf(" %d  ", d.parent[i])
	}

	fmt.Printf("\n --- Disjoint Set- Rank list \n")
	for i := 0; i < d.n; i++ {
		fmt.Printf(" %d  ", d.rank[i])
	}
}

/*
func main() {
	var dj DisjointSet
	dj.init(10)
	dj.WeightedUnion(1, 2)
	dj.Print()
}
*/
