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

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	//////////////////////////////////////
	// n := int32(nTemp)
	cnt := 0
	max := 0
	bi := strconv.FormatInt(nTemp, 2)
	for _, s := range bi {
		if string(s) == "1" {
			cnt++
		} else {
			if cnt > max {
				max = cnt
			}
			cnt = 0
		}

		if cnt > max {
			max = cnt
		}
	}
	fmt.Println(max)
	////////////////////////////////
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
