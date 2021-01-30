package main

import "fmt"

type edge struct { // cesta
	where int
	time  int
}

func main() {
	var n, m, s, t int
	var maze [][]edge
	fmt.Scan(&n, &m, &s, &t) // načtení dat
	maze = make([][]edge, n)
	for i := range maze { // inicializace bludiště
		maze[i] = make([]edge, 0)
	}
	for i := 0; i < m; i++ { // načtení cest
		var a, b, w int
		fmt.Scan(&a, &b, &w)
		maze[a] = append(maze[a], edge{b, w})
	}
}
