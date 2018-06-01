package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	exists := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, exists)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\v", err)
				continue
			}
			countLines(f, counts, exists)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\tfile:%s\n", n, line, exists[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, exists map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if _, ok := counts[text]; ok {
			counts[text]++
			exists[text] = f.Name()
		} else {
			counts[text] = 1
		}
	}
}
