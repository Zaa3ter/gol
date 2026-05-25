package core

// 1 Any live cell with fewer than two live neighbours dies, as if by underpopulation.
// 2 Any live cell with two or three live neighbours lives on to the next generation.
// 3 Any live cell with more than three live neighbours dies, as if by overpopulation.
// 4 Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

type State [][]bool

func (st State) InBounds(p Point) bool {
	if p.X < 0 || p.X >= len(st[0]) || p.Y < 0 || p.Y >= len(st) {
		return false
	}
	return true
}

func (st State) NextRound() {
	newState := CreateState(len(st), len(st[0]))
	for y, row := range st {
		for x := range row {
			p := Point{X: x, Y: y}
			ncount := countNeigbours(st, p)
			if ncount == 2 {
				newState.Set(p, st.Get(p))
			} else {
				newState.Set(p, !st.Get(p))
			}
		}
	}
}

func (st State) Set(p Point, v bool) {
	st[p.Y][p.X] = v
}

func (st State) Toggle(p Point) {
	st.Set(p, !st.Get(p))
}

func (st State) Get(p Point) bool {
	return st[p.Y][p.X]
}

type Point struct {
	X, Y int
}

func addPoint(p1, p2 Point) Point {
	return Point{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

func StateExample() State {
	st := CreateState(60, 30)
	st[18][15] = true
	st[19][15] = true
	st[19][14] = true
	st[20][15] = true
	st[20][16] = true
	return st
}

func CreateState(rows, cols int) State {
	st := make([][]bool, rows)
	for r := range rows {
		st[r] = make([]bool, cols)
	}
	return State(st)
}

var directions = [8]Point{
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
	{-1, 0},
	{-1, 1},
}

func countNeigbours(st State, p Point) (count int) {
	for _, dir := range directions {
		n := addPoint(p, dir)
		if !st.InBounds(n) {
			continue
		}
		if st.Get(p) {
			count++
		}
	}
	return count
}
