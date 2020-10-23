package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getHourGlassValue(arr *[][]int32, i int, j int) *int64 {
	var total int64 = 0
	total += int64((*arr)[i][j])
	total += int64((*arr)[i][j+1])
	total += int64((*arr)[i][j+2])
	total += int64((*arr)[i+1][j+1])
	total += int64((*arr)[i+2][j])
	total += int64((*arr)[i+2][j+1])
	total += int64((*arr)[i+2][j+2])

	return &total
}

func hourGlass(arr *[][]int32) {
	var best int64
	for i := 0; i+2 < len(*arr); i++ {
		for j := 0; j+2 < len((*arr)[i]); j++ {
			total := getHourGlassValue(arr, i, j)

			if *total > best {
				best = *total
			} else if i == 0 && j == 0 {
				best = *total
			}
		}
	}
	fmt.Printf("%v\n", best)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}
	hourGlass(&arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
