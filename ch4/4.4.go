package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	fmt.Println(rotate(s, 3))
}

func rotate(s []int, n int) []int {
	l := len(s)
	n = n % l
	r := make([]int, l)

	for i, j := 0, n; i < l; i, j = i+1, (j+1)%l {
		r[i] = s[j]
	}

	return r
}
