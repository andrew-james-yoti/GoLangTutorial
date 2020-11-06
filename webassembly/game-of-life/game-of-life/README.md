# Game Of Life

[Pseudo Code](https://rosettacode.org/wiki/Conway%27s_Game_of_Life)

module `andrew.james.yoti.com/gameoflife`

## Webassemby

Build the web assembly exec

```shell script
GOOS=js GOARCH=wasm go build -o main.wasm
```

Copy the compiled code to the game-of-life-webserver

```shell script
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ../game-of-life-server/public/javascripts/
```

Copy to the root until I find out how to configure the path
```shell script
cp main.wasm  ../game-of-life-server/public
```

## Game Of Life Server

Need to set the mime type in express app.js

`express.static.mime.define({'application/wasm': ['wasm']});`


## Problems

- `can't load package: package .: build constraints exclude all Go files in /Users/~/GoLangTutorial/webassembly/game-of-life/game-of-life` doesn't appear to cause any build issues (cannot use `go run`), so just appears to be related to VS Code

## Rules

// Rule 1: Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.

(Cell::Alive, x) if x < 2 => Cell::Dead,

// Rule 2: Any live cell with two or three live neighbours lives on to the next generation.

(Cell::Alive, 2) | (Cell::Alive, 3) => Cell::Alive,

// Rule 3: Any live cell with more than three live neighbours dies, as if by overpopulation.

(Cell::Alive, x) if x > 3 => Cell::Dead,

// Rule 4: Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

(Cell::Dead, 3) => Cell::Alive,

// All other cells remain in the same state.

(otherwise, _) => otherwise,