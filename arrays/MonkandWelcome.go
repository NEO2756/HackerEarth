package main

import (
	"fmt"
)

func main() {
  var N, c int
  fmt.Scanf("%d", &N)

  arr := make([]int, N)
  for i := 0; i < N; i++ {
    fmt.Scanf("%d", &arr[i])
  }
  for i := 0; i < N; i++ {
      fmt.Scanf("%d", &c)
      fmt.Printf("%d ", c + arr[i])
  }
  fmt.Println();
}
