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
	maze   [][]edge
	before []place
	food   int
}
type retStat struct {
	end  int
	time int
}

func solve(work Wmaze, decide bool, time, start int) (end, t int) {
	var ret []retStat
	ret = make([]retStat, len(work.maze[start]))
	for i, next := range work.maze[start] {
		tS, tT := solve(work, !decide, time+next.time, next.where)
		ret[i] = retStat{tS, tT}
	}
	for _, r := range ret {
		if decide == LEFT {
			var end bool
			var max int
			switch r.end {
			case RUN:
				return RUN, -1
			case END:
				end = true
			case FOOD:
				if r.time > max {
					max = r.time
				}
			}
			if end {
				return END, -1
			} else {
				return FOOD, max
			}
		} else { // RIGHT
			var min int
			var run, end bool
			switch r.end {
			case RUN:
				run = true
			case END:
				end = true
			case FOOD:
				if min == 0 || min > r.time {
					min = r.time
				}
			}
			if min > 0 {
				return FOOD, min
			} else if run {
				return RUN, -1
			} else if end {
				return END, -1
			}
		}
	}
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
	work = Wmaze{maze, before, t}
	for i := range maze { // inicializace bludiště
		maze[i] = make([]edge, 0)
	}
	for i := 0; i < m; i++ { // načtení cest
		var a, b, w int
		fmt.Scan(&a, &b, &w)
		maze[a] = append(maze[a], edge{b, w})
	}
	end, time := solve(work, LEFT, 0, s)
	switch end {
	case FOOD:
		fmt.Println(time)
	case RUN:
		fmt.Println("maraton")
	case END:
		fmt.Println("uveznen")
	}
}
