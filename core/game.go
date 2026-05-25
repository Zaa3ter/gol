package core

import (
	"math/rand"
	"time"
)

// 1 Any live cell with fewer than two live neighbours dies, as if by underpopulation.
// 2 Any live cell with two or three live neighbours lives on to the next generation.
// 3 Any live cell with more than three live neighbours dies, as if by overpopulation.
// 4 Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

type State struct {
	Cells  []bool
	Rows   int
	Cols   int
	Buffer []bool // Reuse buffer for next generation
}

func (st *State) InBounds(p Point) bool {
	return p.X >= 0 && p.X < st.Cols && p.Y >= 0 && p.Y < st.Rows
}

func (st *State) NextRound() {
	if len(st.Buffer) != len(st.Cells) {
		st.Buffer = make([]bool, len(st.Cells))
	}

	cols := st.Cols
	rows := st.Rows

	for y := range rows {
		for x := range cols {
			ncount := 0
			// Neighbor check
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dy == 0 && dx == 0 {
						continue
					}
					nx, ny := x+dx, y+dy
					if nx >= 0 && nx < cols && ny >= 0 && ny < rows {
						if st.Cells[ny*cols+nx] {
							ncount++
						}
					}
				}
			}

			idx := y*cols + x
			if st.Cells[idx] {
				st.Buffer[idx] = ncount == 2 || ncount == 3
			} else {
				st.Buffer[idx] = ncount == 3
			}
		}
	}
	// Swap buffers
	st.Cells, st.Buffer = st.Buffer, st.Cells
}

func (st *State) Seed() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range st.Cells {
		st.Cells[i] = r.Float32() < 0.2
	}
}

func (st *State) Toggle(p Point) {
	if st.InBounds(p) {
		idx := p.Y*st.Cols + p.X
		st.Cells[idx] = !st.Cells[idx]
	}
}

func (st *State) Get(p Point) bool {
	if !st.InBounds(p) {
		return false
	}
	return st.Cells[p.Y*st.Cols+p.X]
}

type Point struct {
	X, Y int
}

func CreateState(rows, cols int) State {
	return State{
		Cells:  make([]bool, rows*cols),
		Rows:   rows,
		Cols:   cols,
		Buffer: make([]bool, rows*cols),
	}
}

func (st *State) countNeigbours(p Point) (count int) {
	cols := st.Cols
	rows := st.Rows
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}
			nx, ny := p.X+dx, p.Y+dy
			if nx >= 0 && nx < cols && ny >= 0 && ny < rows {
				if st.Cells[ny*cols+nx] {
					count++
				}
			}
		}
	}
	return count
}
