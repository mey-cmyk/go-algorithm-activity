package main

import "fmt"

const INF = 999999

func dijkstra(graph [][]int, start int) {
	n := len(graph)

	dist := make([]int, n)
	visited := make([]bool, n)

	// Awalnya semua jarak = tak hingga
	for i := 0; i < n; i++ {
		dist[i] = INF
	}

	// Jarak dari titik awal ke dirinya sendiri = 0
	dist[start] = 0

	for count := 0; count < n-1; count++ {

		// Cari vertex yang belum dikunjungi
		// dengan jarak paling kecil
		minDist := INF
		current := -1

		for i := 0; i < n; i++ {
			if !visited[i] && dist[i] < minDist {
				minDist = dist[i]
				current = i
			}
		}

		// Tandai vertex sebagai sudah dikunjungi
		visited[current] = true

		// Update jarak ke vertex tetangga
		for i := 0; i < n; i++ {
			if graph[current][i] != 0 && !visited[i] {

				newDist := dist[current] + graph[current][i]

				if newDist < dist[i] {
					dist[i] = newDist
				}
			}
		}
	}

	fmt.Println("Shortest distance from vertex", start)

	for i := 0; i < n; i++ {
		if dist[i] == INF {
			fmt.Println("To", i, "= unreachable")
		} else {
			fmt.Println("To", i, "=", dist[i])
		}
	}
}

func main() {

	// Vertex:
	// 0 = A
	// 1 = B
	// 2 = C
	// 3 = D
	// 4 = E

	graph := [][]int{
		{0, 4, 2, 0, 0},
		{4, 0, 1, 5, 0},
		{2, 1, 0, 2, 4},
		{0, 5, 2, 0, 3},
		{0, 0, 4, 3, 0},
	}

	start := 0

	dijkstra(graph, start)
}