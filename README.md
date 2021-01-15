## Getting Started

Make dir

```shell script
mkdir my-mod
```

Init module

```shell script
go mod init example.com/mymod
```

Build

```shell script
go build
```

Run
```shell script
go run main.go
```

Install package
```shell script
go get github.com/onsi/ginkgo
```

Install module

```shell script
go install example.com/mymod
```

## Creating modules

exported functions must start with Caps e.g. `func HelloWorld() {}`
private functions must start with lowercase e.g. `func helloWorld() {}`