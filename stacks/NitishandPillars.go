/*
Nitish is a short hieghted person. He is standing facing N pillars of different heights with ith pillar having height hi. He tries to see all the possible pillars. He wants to know that how many buildings will he be able to see in the range [L, R] both inclusive.

Input

The first line contains an integer N denoting the number of pillars.
Next line contains N integers denoting height of ith pillar.
Next line contains a single integer Q.
Next Q lines contain pairs L and R respectively.


Output

For every Q queries print the number of buildings visible in the range [L, R].



Constraints

1<=N, Q<=105

1<=L<=R<=N

1<=h<=109
*/

package main

import (
  "fmt"
)

func main() {
  var N int

  arr := make([]int, N)
  result := make([][]int, N*N)

  for i := 0; i < N; i++ {
    fmt.Scan(&arr[i])
  }

  for i := 0; i < N; i++ {
    for j := i; j < N; j++ {
       if arr[i] < arr[j] {
         result[i][j] = j - i - 1
       } else {
         result[i][j] = j - i +
       }
      }
  }
  fmt.Scan(&Q)

  for i := 0; i < Q; i++ {
    fmt.Scan(&start)
    fmt.Scan(&end)

    for idx := start; idx <= end; start++ {

    }
  }
}
