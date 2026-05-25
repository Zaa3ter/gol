package core

import "testing"

func TestCountNeighbours(t *testing.T) {
	st := CreateState(3, 3)
	// Set up a blinker middle
	st.Set(Point{X: 0, Y: 1}, true)
	st.Set(Point{X: 1, Y: 1}, true)
	st.Set(Point{X: 2, Y: 1}, true)

	// Test middle cell (1,1)
	count := st.countNeigbours(Point{X: 1, Y: 1})
	if count != 2 {
		t.Errorf("Expected 2 neighbours for (1,1), got %d", count)
	}
}

func TestNextRound(t *testing.T) {
	st := CreateState(3, 3)
	// Blinker (horizontal)
	st.Set(Point{X: 0, Y: 1}, true)
	st.Set(Point{X: 1, Y: 1}, true)
	st.Set(Point{X: 2, Y: 1}, true)

	st.NextRound()

	// After one round, it should be a vertical blinker
	// (1,0), (1,1), (1,2) should be true
	if !st.Get(Point{X: 1, Y: 0}) || !st.Get(Point{X: 1, Y: 1}) || !st.Get(Point{X: 1, Y: 2}) {
		t.Errorf("NextRound did not produce expected vertical blinker.")
	}
}
