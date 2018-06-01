package main

import (
	"fmt"
	"sort"
	"strings"
)

func compareString(s1, s2 string) bool  {
	arrStr1 := strings.Split(s1, "")
	arrStr2 := strings.Split(s2, "")

	sort.Strings(arrStr1)
	sort.Strings(arrStr2)

	return strings.Join(arrStr1, "") == strings.Join(arrStr2, "")
}

func main()  {
	fmt.Println(compareString("中文1233", "23中文13"))
}
