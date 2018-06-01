package main

import (
	"./popcount"

	"testing"
	"fmt"
)

var testNum uint64 = 12345677892432

func BenchmarkFunction4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = popcount.PopCount4(testNum)
	}
}

func main() {

	br4 := testing.Benchmark(BenchmarkFunction4)
	fmt.Println(br4)
}
