package universe

import (
	"fmt"
	"math/rand"
	"time"
)

// Cell basic structual unit in Universe
type Cell struct {
	Alive bool
}

// Universe where life exists
type Universe struct {
	Cells [10][10]Cell
}

// RandCellValue returns a true / false
func randCellValue() bool {
	rand.Seed(time.Now().Unix())
	if rand.Float32() > 0.5 {
		fmt.Println("rand.Float32() %h", rand.Float32())
		return true
	}

	return false
}

// CreateUniverse initialise cells in universe
func createUniverse() *Universe {
	universe := Universe{}

	for i := 0; i < len(universe.Cells); i++ {
		for j := 0; j < len(universe.Cells[i]); j++ {
			isAlive := randCellValue()
			universe.Cells[i][j].Alive = isAlive
		}
	}

	return &universe
}

// GetRowStartIndex - get the -1 row index
func getRowStartIndex(row int16) int16 {
	if row > 0 {
		return row - 1
	}
	return 0
}

// GetRowEndIndex - get the + 1 row index
func getRowEndIndex(cells *[10][10]Cell, row int16) int16 {
	if row+1 > int16(len(*cells)) {
		return row
	}
	return row + 1
}

// GetColStartIndex - get the -1 column index
func getColStartIndex(col int16) int16 {
	if col > 0 {
		return col - 1
	}
	return 0
}

// GetColEndIndex - get the + 1 column
func getColEndIndex(cells *[10]Cell, col int16) int16 {
	if col+1 > int16(len(*cells)) {
		return col
	}
	return col + 1
}

// CountAdjacentLiveCells - count how may adjacent cells are alive
func countAdjacentLiveCells(cells *[10][10]Cell, row int16, col int16) int8 {
	var liveCells int8 = 0

	for i := getRowStartIndex(row); i <= getRowEndIndex(cells, row); i++ {
		for j := getColStartIndex(col); j <= getColEndIndex(&cells[0], col); j++ {
			if cells[i][j].Alive == true {
				if i != row || j != col {
					liveCells++
				}
			}
		}
	}

	return liveCells
}

// Rule 1: Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
func golRule1(cells *[10][10]Cell, row int16, col int16) {
	liveCells := countAdjacentLiveCells(cells, row, col)

	if liveCells < 2 {
		cells[row][col].Alive = false
	}
}

// Rule 2: Any live cell with two or three live neighbours lives on to the next generation.
func golRule2(universe *Universe, row int16, col int16) {

}

// Rule 3: Any live cell with more than three live neighbours dies, as if by overpopulation.
func golRule3(universe *Universe, row int16, col int16) {

}

// Rule 4: Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
func golRule4(universe *Universe, row int16, col int16) {

}

// GaemOfLife recursive routine
func gameOfLife() {
	// just pause it
	time.Sleep(300 * time.Millisecond)
	// call it again
	gameOfLife()
}
