[Golang Webassembly](https://github.com/golang/go/wiki/WebAssembly#introduction)
[Experiments](https://github.com/stdiopt/gowasm-experiments)

Build the web assembly exec

Set Go operating system to js
Set Go architecture to wasm

```shell script
GOOS=js GOARCH=wasm go build -o main.wasm
```

Copy the compiled code to a webserver

```shell script
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

## Packages
- https://golang.org/pkg/syscall/js/
- https://github.com/dennwc/dom
- https://github.com/agnivade/wasmbrowsertest