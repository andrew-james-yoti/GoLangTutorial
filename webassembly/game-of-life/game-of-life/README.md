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