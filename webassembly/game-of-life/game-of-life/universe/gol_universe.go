package universe

import (
	"fmt"
	"math/rand"
	"time"
)

// Cell basic structual unit in Universe
type Cell struct {
	// Dead  int8
	Alive bool
}

// Universe where life exists
type Universe struct {
	Cells [10][10]Cell
}

// RandCellValue returns a true / false
func RandCellValue() bool {
	rand.Seed(time.Now().Unix())
	if rand.Float32() > 0.5 {
		fmt.Println("rand.Float32() %h", rand.Float32())
		return true
	}

	return false
}

// CreateUniverse initialise cells in universe
func CreateUniverse() *Universe {
	universe := Universe{}

	// var i int = 0
	// for ; i < len(universe.Cells); i++ {
	// 	for j := 0; j < len(universe.Cells[i]); j++ {
	// 		universe.Cells[i][j].Alive = false
	// 	}
	// }

	// return i
	for i := 0; i < len(universe.Cells); i++ {
		for j := 0; j < len(universe.Cells[i]); j++ {
			isAlive := RandCellValue()
			universe.Cells[i][j].Alive = isAlive
		}
	}

	return &universe
}
