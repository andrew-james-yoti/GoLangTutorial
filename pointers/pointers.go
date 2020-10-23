package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "strconv"
)

func formatInput(input *string) *[]string {
	fmt.Printf("%v", *input)
	strArr := strings.Split(*input, " ")
	return &strArr
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	input := scanner.Text()

	strArr := formatInput(&input)

	for i := 0; i < len(*strArr); i++ {
		fmt.Printf("%v\n", (*strArr)[i])
	}
}
