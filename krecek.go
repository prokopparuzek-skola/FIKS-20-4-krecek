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
	LATE
)
const ( // levá a pravá hemisféra, index v poli before
	L = 0
	R = 1
)

type edge struct { // cesta
	where int
	time  int
}

type place struct {
	there bool
	when  int
}

type Wmaze struct { // bludiště + seznam kde už křeček byl
	maze   [][]edge
	before [][2]place
	food   int
}
type retStat struct {
	end  int
	time int
}

func solve(work Wmaze, decide bool, time, start int) (int, int) {
	// nají se?
	if start == work.food {
		return FOOD, time
	}
	// už tu byl...
	var who int
	if decide == LEFT {
		who = L
	} else {
		who = R
	}
	if work.before[start][who].there {
		return RUN, -1
	} else if work.before[start][who].when > 0 && time >= work.before[start][who].when { // byl tu v jiné větvi dřív
		return LATE, -1
	}
	work.before[start][who].when = time
	work.before[start][who].there = true // byl jsem tu
	// ponoření
	var ret []retStat
	ret = make([]retStat, len(work.maze[start]))
	for i, next := range work.maze[start] { // projdi všechny cesty co následují
		tS, tT := solve(work, !decide, time+next.time, next.where)
		ret[i] = retStat{tS, tT}
	}
	work.before[start][who].there = false // uvolnění pro průchody z jiných křižovatek
	var max, min int
	var run, end bool
	if decide == LEFT { // LEFT
		for _, r := range ret {
			switch r.end {
			case RUN: // můžu běžet maraton
				return RUN, -1
			case END: // nebo se uvěznit
				end = true
			case FOOD: // či najdu jídlo
				if r.time > max { // před jídlem chci běžet co nejdéle
					max = r.time
				}
			}
		}
		if end { // nejde maraton, tak se křeček uvězní
			return END, -1
		} else if max != 0 { // když musí, tak jde k jídlu
			return FOOD, max
		} else {
			return LATE, -1
		}
	} else { // RIGHT
		for _, r := range ret {
			switch r.end {
			case RUN: // běh
				run = true
			case END: // vězení
				end = true
			case FOOD: // jídlo
				if min == 0 || min > r.time { // chci nejkratší(časově) cestu
					min = r.time
				}
			}
		}
		if min > 0 { // k jídlu
			return FOOD, min
		} else if run { // běž
			return RUN, -1
		} else if end { // vězení
			return END, -1
		} else {
			return LATE, -1
		}
	}
}

func main() {
	var n, m, s, t int
	var maze [][]edge
	var before [][2]place
	var work Wmaze

	fmt.Scan(&n, &m, &s, &t) // načtení dat
	maze = make([][]edge, n)
	before = make([][2]place, n)
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
