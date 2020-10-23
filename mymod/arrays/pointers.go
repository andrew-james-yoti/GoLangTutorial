// Package arrays does some stuff with arrays and pointers
package arrays

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func formatInput(input *string) *[]string {
	// fmt.Printf("%v", *input)
	strArr := strings.Split(*input, " ")
	return &strArr
}

// TextPointers doesn't return anything, it just does what it does
func TextPointers() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	input := scanner.Text()

	strArr := formatInput(&input)

	fmt.Printf("\n")
	for i := 0; i < len(*strArr); i++ {
		fmt.Printf("%v\n", (*strArr)[i])
	}
}
