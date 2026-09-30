package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	from   int
	to     int
	weight int
}

type DSU struct {
	parent []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &DSU{
		parent: parent,
	}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}

	return d.parent[x]
}

func (d *DSU) Union(a int, b int) bool {
	rootA := d.Find(a)
	rootB := d.Find(b)

	if rootA == rootB {
		return false
	}

	d.parent[rootB] = rootA

	return true
}

func kruskal(edges []Edge, vertices int) {

	// Urutkan edge berdasarkan weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	dsu := NewDSU(vertices)

	totalWeight := 0
	edgeCount := 0

	fmt.Println("Edges selected:")

	for _, edge := range edges {

		// Masukkan edge hanya jika
		// tidak membuat cycle
		if dsu.Union(edge.from, edge.to) {

			fmt.Println(
				edge.from,
				"-",
				edge.to,
				"weight:",
				edge.weight,
			)

			totalWeight += edge.weight
			edgeCount++

			if edgeCount == vertices-1 {
				break
			}
		}
	}

	fmt.Println("Total weight:", totalWeight)
}

func main() {

	// Vertex:
	// 0 = A
	// 1 = B
	// 2 = C
	// 3 = D
	// 4 = E

	edges := []Edge{

		{0, 2, 3},  // A-C
		{2, 1, 10}, // C-B
		{2, 3, 2},  // C-D
		{2, 4, 4},  // C-E
		{1, 3, 6},  // B-D
		{3, 4, 1},  // D-E
	}

	kruskal(edges, 5)
}