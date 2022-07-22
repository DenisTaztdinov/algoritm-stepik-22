package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	a1 := []rune(a)
	a2 := []rune("")
	for j := 0; j < len(a1); j++ {
		count := strings.Count(a, string(a1[j]))
		if count == 1 {
			a2 = append(a2, a1[j])
		}
	}
	fmt.Println(string(a2))
}
