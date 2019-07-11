/*
Given an array  of size , you can jump from an index  to another index  if  >= , for  > . Find the length of the longest sequence of jumps that can be possible in the array. You can start at any index.

Input Format:

First line contains an integer .

Second line contains the integer .

Third line contains  space separated integers (The array )

Output Format:

Print the required length.

Constrains:

 ≤  ≤

*/

package main

import (
	"fmt"
)

//var arr, v, q, sum []int


func main() {

  var K, N int

  fmt.Scan(&K)
  fmt.Scan(&N)

  arr := make([]int, N)
	dp := make([]int, N)
	max := -1

	for i := 0; i < N; i++ {
     fmt.Scan(&arr[i])
  }

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if arr[j] - arr[i] >= K {
				if dp[i] + 1 > dp[j] {
					dp[j] = dp[i] + 1
					if dp[j] > max {
						max = dp[j]
					}
				}
			}
		}
	}
	if max == -1 {
		fmt.Println(0)
	} else {
		fmt.Println(max+1)
	}
}
