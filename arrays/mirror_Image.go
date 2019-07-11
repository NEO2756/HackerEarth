/*
You are given a binary tree rooted at 1. You have to find the mirror image of any node qi about node 1. If it doesn't exist then print -1.
Input:
First line of input is N and Q.
Next N-1 line consists of two integers and one character first of whose is parent node , second is child node and character "L" representing Left child and "R" representing right child.
Next Q lines represents qi.
Output:
For each qi print it mirror node if it exists else print -1.
NOTE: 1 is mirror image of itself.
Constraints:
1 <= N <= 103
1<= Q <= 103

*/

package main

import (
	"fmt"
)
var nodes, mirror []int

func image (left, right, given int) (int) {

    //fmt.Println(left, right, given)
    if left == 0 && right == 0 {
      return -1
    }
    if left != 0 && right != 0 {
      mirror[left] = right
      mirror[right] = left
    }
    if left == given {
       return mirror[left]
    }
    if right == given {
      return mirror[right]
    }

    if nodes[2*left] > 0 && nodes[2*right + 1] > 0 {
      if v := image(nodes[2*left], nodes[2*right + 1], given); v != -1 {
        return v
      }
    }
    if nodes[2*left+1] > 0 &&  nodes[2*right] > 0 {
      if v := image(nodes[2*left+1], nodes[2*right], given); v != -1 {
        return v
      }
    }
    return -1
}

func main() {

  var N, Q, src, dest, q int
  var direction rune

  fmt.Scan(&N)
  fmt.Scan(&Q)

  nodes = make([]int, N*2 + 2) //N+1 nodes
  mirror = make([]int, N*2 + 2)

  mirror[1] = 1  //1 is morror to itself
  nodes[1] = 1
  for i := 1; i < N; i++ {
    fmt.Scan(&src)
    fmt.Scan(&dest)
    fmt.Scanf("%c", &direction)
    //fmt.Println(src, dest, string(direction))
    if direction == 'R' {
      nodes[2 * src + 1] = dest
    } else {
      nodes[2 * src] = dest
    }
  }

  for i := 1 ; i <= Q; i++ {
    fmt.Scan(&q)
    if q == 1 {
      fmt.Println(q)
      continue
    }
    ans := image(nodes[2 * 1], nodes[2*1 + 1], q)
    if ans == 0 {
      ans = -1
    }
    fmt.Println(ans)
  }

  //fmt.Println(image(nodes[2 * 1], nodes[2*1 + 1], 2))


}
