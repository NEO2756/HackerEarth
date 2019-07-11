package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var string string
	fmt.Scanf("%s", &string)
	if strings.Compare(string, reverse(string)) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
