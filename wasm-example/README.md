```shell script
GOOS=js GOARCH=wasm go build -o main.wasm
```

Copy the compiled code to a webserver

```shell script
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```