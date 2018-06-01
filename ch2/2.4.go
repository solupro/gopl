package main

import (
	"./popcount"

	"testing"
	"fmt"
)

var testNum uint64 = 12345677892432

func BenchmarkFunction3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = popcount.PopCount3(testNum)
	}
}

func main() {

	br3 := testing.Benchmark(BenchmarkFunction3)
	fmt.Println(br3)
}
