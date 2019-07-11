package main

import (
	"fmt"
	//"strings"
)

func main() {

  var N, Q ,q, l, r int
  fmt.Scanf("%d %d", &N, &Q)

  s := make([]int, N)
  for i := 0; i < N; i++ {
    fmt.Scanf("%d", &s[i])
  }
  for i := 0; i < Q; i++ {
    fmt.Scanf("%d", &q)
    if q == 0 {
      fmt.Scanf("%d %d", &l, &r)
      if s[r-1] == 0 {
        fmt.Println("EVEN")
      } else {
        fmt.Println("ODD")
      }
    } else {
      fmt.Scanf("%d", &l)
      s[l - 1] ^= 1
    }
  }
}
