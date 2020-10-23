package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func spaceSeperateStrings(text string) {
	var strArr []string = strings.Split(text, "")
	var oddIndexString string
	var evenIndexString string

	for i := 0; i < len(strArr); i++ {
		if i > 0 {
			if i%2 == 0 {
				evenIndexString += strArr[i]
			} else {
				oddIndexString += strArr[i]
			}
		} else {
			evenIndexString = strArr[0]
		}
	}
	fmt.Printf("%v %v\n", evenIndexString, oddIndexString)
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	testCases, _ := strconv.ParseInt(scanner.Text(), 10, 32)

	var i int32 = 0

	for ; i < int32(testCases); i++ {
		scanner.Scan()
		spaceSeperateStrings(scanner.Text())
	}
}
