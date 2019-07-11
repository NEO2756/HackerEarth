/*


There is a complete Binary tree in which there are N nodes. Each vertex of the tree is assigned a value.

You are given the values of all nodes in level order traversal of the binary tree (from left to right).

You are given three types of queries.

1 X  LV  Find the value of Xth node from left on the level LV.

2  L R  Find the the sum of values of nodes from level L to R (L and R are inclusive).

3  X  LV  VAL Update the value of Xth node from left which is at level LV to VAL.

Note: Root is at level 1.

INPUT

In the first line you are given two integers N, Q.

In the next line you are given an array A where the ith value denotes the value of ith node.

In the next Q lines you are given one of the three queries.

It is guaranteed that there exists at least one query of type 1 or 2.

Also it is guaranteed all given queries are valid.

OUTPUT

For each query of 1st and 2nd type print the required value.


*/

package main

import (
  "fmt"
  //"math"
)

func main() {


  //var N, Q int
  var N, Q, node, level, val, query, start, end int
  fmt.Scan(&N)
  fmt.Scan(&Q)

  arr := make([]int, N+1)
  sum := make([]int, N+1)
  for i := 1; i <= N; i++ {
    fmt.Scan(&arr[i])
  }


  temp := N
  for l := 1; temp > 0; l++ {  //sum at each level
    for i := 1 << uint(l - 1); i <= 2 * 1 << uint(l - 1) - 1 && i <= N; i++ {
      sum[l] += arr[i]
      //fmt.Printf("%d", i)
      temp--
    }
    //fmt.Printf("\n")
    //fmt.Println("Level", l , "sum = ", sum[l])
  }

  for q := 1; q <= Q; q++ {

    fmt.Scan(&query)
    if query == 1 {
      fmt.Scan(&node)
      fmt.Scan(&level)

      index := (1 << uint(level - 1)) + (node - 1)
      fmt.Println(arr[index])
    } else if query == 2 {
      sum1 := 0
      fmt.Scan(&start)
      fmt.Scan(&end)

      //fmt.Println("here", query, sum, start, end)
      for i := start; i <= end && i <= N; i++ {
          sum1 += sum[i]
      }
      fmt.Println(sum1)
    } else {
      fmt.Scan(&node)
      fmt.Scan(&level)
      fmt.Scan(&val)

      index := (1 << uint(level - 1)) + (node - 1)
      sum[level] -= arr[index]
      arr[index] = val
      sum[level] += val
      //fmt.Println(sum)
      //fmt.Println(index, 1 << uint(level - 1) )
    }
  }
}
