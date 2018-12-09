package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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

	max := int32(0)
	sum := int32(0)
	first := true
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			sum = 0
			sum += arr[x][y] + arr[x][y+1] + arr[x][y+2]
			sum += arr[x+1][y+1]
			sum += arr[x+2][y] + arr[x+2][y+1] + arr[x+2][y+2]

			if first {
				max = sum
				first = false
				continue
			}

			if sum > max {
				max = sum
			}
		}
	}
	fmt.Println(max)
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
