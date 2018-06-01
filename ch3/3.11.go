package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer

	if dot := strings.LastIndex(s, "."); dot > 0 {
		buf = doComma(s[:dot])
		buf.WriteString(s[dot:])
	} else {
		buf = doComma(s)
	}

	return buf.String()
}

func doComma(s string) bytes.Buffer {
	var buf bytes.Buffer
	var idx int
	k := 0
	for i := len(s) - 1; i >= 0; i-- {
		buf.WriteByte(s[i])
		idx = k + 1
		if idx % 3 == 0 && idx != 0 {
			buf.WriteByte(',')
		}
		k++
	}

	s = buf.String()
	buf.Reset()
	for i := len(s) - 1; i >= 0; i-- {
		buf.WriteByte(s[i])
	}

	return buf
}

func main()  {
	fmt.Println(comma("3456789.12345"))
}