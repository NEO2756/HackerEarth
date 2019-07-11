package main

import (
	"fmt"
	//"strings"
)

func main() {

  var T, N, sum, j int

  fmt.Scanf("%d", &T)

  for i := 0; i < 2* T; i += 2 {
    sum, j = 0, 0
    fmt.Scanf("%d", &N)
    arr := make([]int, N)
    for j := 0 ; j < N; j++ {
      fmt.Scanf("%d", &arr[j])
    }
    //fix the start
    m := make(map[int]int)
    for start := 0; start < N; start++ {
      for ; j < N ; j++ {
        _, ok := m[arr[j]]
        if (ok == true) {
          break
        }
        //fmt.Println("insert ", arr[j])
        m[arr[j]] = arr[j]
      }
      sum += ((j - start) * (j - start + 1))/ 2
      //fmt.Println("sum now = ",sum, "start, j ", start, j)
      delete(m, arr[start])
    }
    fmt.Println(sum)
  }

}
