package main

import (
	"fmt"
)

func main() {

	var N, c, Q int
	fmt.Scanf("%d", &N)

	arr := make([]int, 100000)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &c)
		arr[c]++
	}

	fmt.Scanf("%d", &Q)
	for i := 0; i < Q; i++ {
		fmt.Scanf("%d", &c)
		if arr[c] > 0 {
			fmt.Println(arr[c])
		} else {
			fmt.Println("NOT PRESENT")
		}
	}
}
