package main

import (
	"fmt"

	"example.com/user/mymod/arrays"
	"example.com/user/mymod/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, Go!"))
	arrays.TextPointers()
}
