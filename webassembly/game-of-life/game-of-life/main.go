// +build js,wasm

package main

import "syscall/js"

func main() {
	document := js.Global().Get("document")
	gameOfLifeCanvas := document.Call("getElementById", "game-of-life-canvas")
	gameOfLifeCtx := gameOfLifeCanvas.Call("getContext", "2d")
	gameOfLifeCtx.Set("fillStyle", "#FFA500")
	gameOfLifeCtx.Call("fillRect", 10, 10, 50, 50)

	// p := document.Call("createElement", "p")
	// p.Set("innerHTML", "Hello WASM from Go!")
	// document.Get("body").Call("appendChild", p)
}
