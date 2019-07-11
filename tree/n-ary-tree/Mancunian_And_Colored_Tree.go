/*


After a hectic week at work, Mancunian and Liverbird decide to go on a fun weekend camping trip. As they were passing through a forest, they stumbled upon a unique tree of N nodes. Vertices are numbered from 1 to N.

Each node of the tree is assigned a color (out of C possible colors). Being bored, they decide to work together (for a change) and test their reasoning skills. The tree is rooted at vertex 1. For each node, they want to find its closest ancestor having the same color.

Input format
The first line contains two integers N and C denoting the number of vertices in the tree and the number of possible colors.
The second line contains  integers. The  integer denotes the parent of the  vertex.
The third line contains N integers, denoting the colors of the vertices. Each color lies between 1 and C inclusive.

Output format
Print N space-separated integers. The  integer is the vertex number of lowest ancestor of the  node which has the same color. If there is no such ancestor, print 1 for that node.
*/


package main

import (
  "fmt"
  //"math"
)

func main() {

  var N, C int

  fmt.Scan(&N)
  fmt.Scan(&C)

  ancestor := make ([]int, N+2)
  color := make([]int, N+2)

  ancestor[1] = -1
  for i := 1; i <= N - 1; i++ {
    fmt.Scan(&ancestor[i+1])
  }
  for i := 1; i <= N; i++ {
    fmt.Scan(&color[i])
  }

  for i := 1; i <= N; i++ {
    targetColor := color[i]
    k := i
    //fmt.Println("k = ", k, "targetColor = ", targetColor)
    for ; ancestor[k] != -1 && targetColor != color[ancestor[k]]; {
      //fmt.Println("ancestor[k] = ", ancestor[k], "color[ancestor[k]] = ", color[ancestor[k]])
      k = ancestor[k]
    }
    fmt.Printf("%d ", ancestor[k])
  }


}
