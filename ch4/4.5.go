package main

import "fmt"

func main() {
	s := []string{"abc", "ka", "ka", "poj"}

	fmt.Println(nonRepeat(s))
}

func nonRepeat(s []string) []string {
	i := 0
	var prev string
	for _, v := range s {
		if v == prev {
			continue
		}

		s[i] = v
		prev = v
		i += 1
	}

	return s[:i]
}
