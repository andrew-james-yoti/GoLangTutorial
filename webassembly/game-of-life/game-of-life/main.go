// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	"andrew.james.yoti.com/gameoflife/game"
)

func main() {
	document := js.Global().Get("document")
	gameOfLifeCanvas := document.Call("getElementById", "game-of-life-canvas")
	gameOfLifeCtx := gameOfLifeCanvas.Call("getContext", "2d")
	gameOfLifeCtx.Set("fillStyle", "#FFA500")
	gameOfLifeCtx.Call("fillRect", 10, 10, 50, 50)

	gameOfLifeCtx.Set("fillStyle", "rgba(0, 0, 200, 0.5)")
	gameOfLifeCtx.Call("fillRect", 30, 30, 50, 50)

	fmt.Println("%d", game.TestFunc())
	// p := document.Call("createElement", "p")
	// p.Set("innerHTML", "Hello WASM from Go!")
	// document.Get("body").Call("appendChild", p)
}
