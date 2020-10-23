/*
10499* go build hello.go
10500* ./hello
10501  go build hello.go
10502  ./hello
10503  go build hello.go
10504  ./hello
10505* go mod init example.com/user/hello
10506* cat go.mod
10507* go install example.com/user/hello
10508* export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
10509* hello
10510  cd morestrings
10511  go build
10512  go build morestrings.go
10513  cd ../
10514  go install example.com/user/hello
10515  hello
10516  export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
10517  hello
10518  go install example.com/user/hello
10519  hello
*/
package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/user/hello/morestrings"
)

func doOperator() {
	var a float32 = 12.1
	var b float32 = 13.2

	sum := a + b

	fmt.Printf("%v\n", sum)
}

func percent() {
	var fraction float64
	fraction = float64(20) / float64(100)
	fmt.Printf("%v\n", fraction)
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	message := "Hello World!\n"
	fmt.Printf(message)
	text, _ := scanner.ReadString('\n')
	fmt.Println(morestrings.ReverseRunes(text))
	doOperator()
	percent()
}
