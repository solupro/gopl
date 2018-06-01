package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	var idx int
	for k, v := range(s) {
		buf.WriteRune(v)
		idx = k + 1
		if idx % 3 == 0 && idx != len(s) {
			buf.WriteByte(',')
		}

	}

	return buf.String()
}

func main()  {
	fmt.Println(comma("123456789"))
}