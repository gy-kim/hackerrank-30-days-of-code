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
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReaderSize(os.Stdin, 21024*1024)
	cnt := int64(0)
	n := int64(0)
	phoneBook := make(map[string]string)
	for {
		cnt++
		bytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		str := strings.TrimRight(string(bytes), "\r\n")

		if cnt == 1 {
			n, _ = strconv.ParseInt(str, 10, 64)
			continue
		}
		if cnt-1 <= n {
			arr := strings.Split(str, " ")
			phoneBook[arr[0]] = arr[1]
			continue
		}

		v, ok := phoneBook[str]
		if !ok {
			fmt.Println("Not found")
			continue
		}
		fmt.Printf("%v=%v\n", str, v)
	}
}
