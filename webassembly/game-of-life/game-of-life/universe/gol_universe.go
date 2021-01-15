package universe

import (
	"math/rand"
	"time"
)

const rows = 20
const columns = 20
const randMax = 30
const randThreshold = 20

// Cell basic structual unit in Universe
type Cell struct {
	Alive bool
}

// Universe where life exists
type Universe struct {
	Cells [rows][columns]Cell
}

// CreateUniverse initialise cells in universe
func CreateUniverse() *Universe {
	universe := Universe{}
	rand.Seed(time.Now().Unix())

	for i := 0; i < len(universe.Cells); i++ {
		for j := 0; j < len(universe.Cells[i]); j++ {

			r := rand.Intn(randMax)
			if r > randThreshold {
				universe.Cells[i][j].Alive = true
			}
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
func getRowEndIndex(cells *[rows][columns]Cell, row int16) int16 {
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
func getColEndIndex(cells *[rows]Cell, col int16) int16 {
	if col+1 > int16(len(*cells)) {
		return col
	}
	return col + 1
}

// CountAdjacentLiveCells - count how may adjacent cells are alive
func countAdjacentLiveCells(cells *[rows][columns]Cell, row int16, col int16) int8 {
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

// GolRule1 Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
func GolRule1(cells *[rows][columns]Cell, row int16, col int16) {
	liveCells := countAdjacentLiveCells(cells, row, col)

	if cells[row][col].Alive && liveCells < 2 {
		cells[row][col].Alive = false
	}
}

// Rule 2: Any live cell with two or three live neighbours lives on to the next generation.
// func GolRule2(cells *[rows][columns]Cell, row int16, col int16) {
// 	liveCells := countAdjacentLiveCells(cells, row, col)
// 	if cells[row][col].Alive {
// 		if liveCells == 2 || liveCells == 3 {
// 			// do nothing, cell stays alive
// 		}
// 	}

// }

// GolRule3 Any live cell with more than three live neighbours dies, as if by overpopulation.
func GolRule3(cells *[rows][columns]Cell, row int16, col int16) {
	liveCells := countAdjacentLiveCells(cells, row, col)

	if cells[row][col].Alive && liveCells > 3 {
		cells[row][col].Alive = false
	}
}

// GolRule4 Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
func GolRule4(cells *[rows][columns]Cell, row int16, col int16) {
	liveCells := countAdjacentLiveCells(cells, row, col)

	if cells[row][col].Alive == false && liveCells == 3 {
		cells[row][col].Alive = true
	}
}
