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

var arr, v, q, sum []int
var start, end int

func qu(i int) {
  q[end] = i
  end++
}

func dq() int {
  e := q[start]
  start++
  return e
}

func empty() bool {
  if start == end {
    return true
  }
  return false
}

func main() {

  var K, N int

  fmt.Scan(&K)
  fmt.Scan(&N)

  arr = make ([]int, N)
  v = make ([]int, N)
  q = make ([]int, N+10)
  sum = make ([]int, N)

  for i := 0; i < N; i++ {
     fmt.Scan(&arr[i])
  }

  qu(0)
  fmt.Println("queued", arr[0], start, end)
  v[0] = 1
  sum[0] = 0
  for empty() == false {
    e := dq()
    fmt.Println("dqueued", arr[e], start, end)
    for j := e + 1; j < N; j++ {
      if arr[j] - arr[e] >= K {
        if sum[e] + 1 > sum[j] {
          sum[j] = sum[e] + 1
        }
        if v[j] == 0 {
          qu(j)
          v[j] = 1
          fmt.Println("queued", arr[j], start, end)
        }
      }
    }
    fmt.Println("sum = ", sum)
  }
 if sum[N-1] != 0 {
   fmt.Println(sum[N-1] + 1)
 } else {
   fmt.Println(sum[N-1])
 }

}
