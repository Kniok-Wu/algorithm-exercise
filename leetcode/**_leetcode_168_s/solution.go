package main

import (
	"fmt"
	"strings"
)

func convertToTitle(columnNumber int) string {
	records := make([]byte, 0)
	for n := 0; columnNumber > 0; n++ {
		cur := (columnNumber-1)%26 + 1
		records = append([]byte{byte(64 + cur)}, records...)
		columnNumber = (columnNumber - cur) / 26
	}

	str := strings.Builder{}
	str.Write(records)
	return str.String()
}

func main() {
	fmt.Println(convertToTitle(28))
}
