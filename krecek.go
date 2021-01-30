package main

import "fmt"

const (
	LEFT  = true  // běh, začíná
	RIGHT = false // jídlo
)
const (
	FOOD = iota
	END
	RUN
)

type edge struct { // cesta
	where int
	time  int
}
type place struct {
	left   bool
	LWhere int
	right  bool
	RWhere int
}
type Wmaze struct {
	maze   *[][]edge
	before *[]place
}

func solve(work *Wmaze, decide bool, time int) (end, t int) {
	return
}

func main() {
	var n, m, s, t int
	var maze [][]edge
	var before []place
	var work Wmaze

	fmt.Scan(&n, &m, &s, &t) // načtení dat
	maze = make([][]edge, n)
	before = make([]place, n)
	work = Wmaze{&maze, &before}
	for i := range maze { // inicializace bludiště
		maze[i] = make([]edge, 0)
	}
	for i := 0; i < m; i++ { // načtení cest
		var a, b, w int
		fmt.Scan(&a, &b, &w)
		maze[a] = append(maze[a], edge{b, w})
	}
	end, time := solve(&work, LEFT, 0)
	switch end {
	case FOOD:
		fmt.Println(time)
	case RUN:
		fmt.Println("maraton")
	case END:
		fmt.Println("uveznen")
	}
}
