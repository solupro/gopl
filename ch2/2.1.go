package main

import (
	"fmt"
	"./conv"
)

func main() {
	v := conv.Kelvin(1.0)
	fmt.Printf("%v = %v\n", v, conv.K2C(v))
}