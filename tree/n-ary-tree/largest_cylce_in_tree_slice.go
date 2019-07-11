/*
You are given a tree of  nodes and  edges. Now you need to select two nodes  and  in the tree such that the cycle that will be formed after adding an edge between the two nodes  and , its length should be maximum. If there are more than one possible answers, you can output any of them.

Input
The first line contains an integer  as input. Next  lines contain a pair of integers  that denote there is an edge between the two nodes  and  in the tree.

Output
In the output, you need to print two integers separated by space which denote the nodes between which you can add the edge so as to maximize the length of the cycle in the tree.

Constraints

*/

package main

import (
  "fmt"
)

var list [][]int
var visited []int
var targetnode, minDepth, targetdepth int

func reset(n int) {
  for i,_ := range visited {
    visited[i] = 0
  }
}

func dfs(depth, node int) {

  //fmt.Println("visited node = ", node, depth)
  visited[node] = 1

  if depth > targetdepth {
    targetnode = node
    targetdepth = depth
  }

  if list[node] == nil {
    return
  }

  s := list[node]
  for _, v := range s {
    if visited[v] == 1 {
      continue
    }
    dfs(depth + 1, v)
  }
}

func main() {

  var N, src, dest int

  fmt.Scan(&N)

  list = make([][]int, N+1)
  visited = make([]int, N+1)

  targetnode, minDepth, targetdepth = 0, 0, 0

  for i := 1; i < N; i++ {
      fmt.Scan(&src)
      fmt.Scan(&dest)
      list[src] = append(list[src], dest)
      list[dest] = append(list[dest], src)
  }
  //fmt.Println(dfs(0, 1))
    dfs(0, 1)
    tmp := targetnode
    reset(N)
    targetnode, targetdepth = 0, 0
    dfs(0, tmp)
    fmt.Println(tmp, targetnode)
}
