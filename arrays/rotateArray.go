package main

import (
	"fmt"
	//"strings"
)

func main() {
	var arr [10000000]int
	var T, N, D, tmp int
	fmt.Scanf("%d", &T)
	for i := 0; i < T ; i++ {
		fmt.Scanf("%d %d", &N, &D)
	//	fmt.Println("N", N,  "D", D)
		for j := 0; j < D; j++ {
			fmt.Scanf("%d", &arr[j])
		}
		for i := D; i < N ; i++ {
			fmt.Scanf("%d", &tmp)
			fmt.Printf("%d ", tmp)
		}
		for j := 0; j < D; j++ {
			fmt.Printf("%d ", arr[j])
		}
		fmt.Println("")
	}
}
