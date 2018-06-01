package main

import (
	"flag"
	"os"
	"fmt"
	"./conv"
)

func main() {
	var t int
	var v float64
	flag.IntVar(&t, "type", 0, "1:C2F,2:F2C,3:Meter2Feet,4:Kilo2Pound")
	flag.Float64Var(&v, "value", 0.0, "press a number")

	flag.Parse()
	if 0 == t {
		flag.Usage()
		os.Exit(0)
	}

	switch t {
	case 1:
		in := conv.Celsius(v)
		fmt.Printf("%v = %v \n", in, conv.C2F(in))
	case 2:
		in := conv.Fahrenheit(v)
		fmt.Printf("%v = %v \n", in, conv.F2C(in))
	case 3:
		fmt.Printf("%v M = %v F\n", v, conv.Meter2Feet(conv.Meter(v)))
	case 4:
		fmt.Printf("%v KG = %v P\n", v, conv.Kilo2Pound(conv.Kilo(v)))
	default:
		fmt.Println("no found type")
	}
}