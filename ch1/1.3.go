package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type joinFunc func() string

func main() {
	cost(normal)()
	cost(fast)()
}

func cost(f joinFunc) func() {
	return func() {
		st := time.Now()
		result := f()
		et := time.Now()
		fmt.Printf("result:%s, cost:%v\n", result, et.Sub(st))
	}
}

func normal() string {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}

	return s
}

func fast() string {
	return strings.Join(os.Args, " ")
}
