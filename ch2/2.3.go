package main

import (
	"./popcount"

	"testing"
	"fmt"
)

var testNum uint64 = 12345677892432

func BenchmarkFunction1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = popcount.PopCount(testNum)
	}
}

func BenchmarkFunction2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = popcount.PopCount2(testNum)
	}
}

func main() {

	br1 := testing.Benchmark(BenchmarkFunction1)
	fmt.Println(br1)

	br2 := testing.Benchmark(BenchmarkFunction2)
	fmt.Println(br2)
}
