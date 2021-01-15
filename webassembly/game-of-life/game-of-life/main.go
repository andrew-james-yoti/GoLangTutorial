// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"
	"time"

	"andrew.james.yoti.com/gameoflife/universe"
)

func gameOfLife(golUniverse *universe.Universe) {
	for i := 0; i < len(&golUniverse.Cells)-1; i++ {
		for j := 0; j < len(&golUniverse.Cells[j])-1; j++ {
			// fmt.Println("Table size rows %d, cols %d", len(&golUniverse.Cells), len(&golUniverse.Cells[j]))
			// fmt.Printf("golUniverse.Cells[%d][%d].Alive - %s", i, j, strconv.FormatBool(golUniverse.Cells[i][j].Alive))
			universe.GolRule1(&golUniverse.Cells, int16(i), int16(j))
			universe.GolRule3(&golUniverse.Cells, int16(i), int16(j))
			universe.GolRule4(&golUniverse.Cells, int16(i), int16(j))
			// fmt.Println("golUniverse.Cells[%d][%d].Alive - %s", i, j, strconv.FormatBool(golUniverse.Cells[i][j].Alive))
			document := js.Global().Get("document")
			tdId := fmt.Sprintf("row%dcolumn%d", i, j)
			td := document.Call("getElementById", tdId)
			if golUniverse.Cells[i][j].Alive == true {
				td.Set("style", "background-color: green;")
			} else {
				td.Set("style", "background-color: none;")
			}
		}
	}
	// just pause it
	time.Sleep(30 * time.Millisecond)
	// call it again
	gameOfLife(golUniverse)
}

func main() {
	document := js.Global().Get("document")
	// gameOfLifeCanvas := document.Call("getElementById", "game-of-life-canvas")
	// gameOfLifeCtx := gameOfLifeCanvas.Call("getContext", "2d")

	gameOfLifeTable := document.Call("createElement", "table")

	golUniverse := universe.CreateUniverse()

	for i := 0; i < len(golUniverse.Cells); i++ {
		tr := document.Call("createElement", "tr")
		trId := fmt.Sprintf("row%d", i)
		tr.Set("id", trId)

		for j := 0; j < len(golUniverse.Cells[i]); j++ {
			td := document.Call("createElement", "td")
			tdId := fmt.Sprintf("row%dcolumn%d", i, j)
			td.Set("id", tdId)
			if golUniverse.Cells[i][j].Alive == true {
				td.Set("style", "background-color: green;")
			}
			tr.Call("appendChild", td)
		}
		gameOfLifeTable.Call("appendChild", tr)
	}
	document.Get("body").Call("appendChild", gameOfLifeTable)
	// gameOfLifeCtx.Set("fillStyle", "#FFA500")
	// gameOfLifeCtx.Call("fillRect", 10, 10, 50, 50)

	// gameOfLifeCtx.Set("fillStyle", "rgba(255, 0, 0, 1)")
	// gameOfLifeCtx.Call("fillRect", 30, 30, 50, 50)

	// p := document.Call("createElement", "p")
	// p.Set("innerHTML", "Hello WASM from Go!")
	// document.Get("body").Call("appendChild", p)
	gameOfLife(golUniverse)
}
