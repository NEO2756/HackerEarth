/*


You are given a directory tree of  directories/folders. Each directory is represented by a particular  which ranges from  to . The id of the root directory is , then it has some child directories, those directories may contain some other ones and it goes on. Now you are given a list of directory id's to delete, you need to find the minimum number of directories that need to be deleted so that all the directories in the given list get deleted.

Note that if you delete a particular directory, all its child directories will also get deleted.

Input
The first line of input contains an integer  that denotes how many folders are there.
The next line contains  space separated integers that where the  integer denotes the id of the parent of the directory with id  . Note that the first integer will be  as  is the id of root folder and it has no parent. Rest of the integers will not be  .
The next line contains an integer  that denotes how many directories you need to delete.
The next line contains  space separated integers that denote the ids of the directories you need to delete.

Output
Print the minimum number of directories that need to be deleted.
*///

package main

import (
  "fmt"
  //"math"
)

func main() {

  var N, C, t int

  fmt.Scan(&N)

  ancestor := make ([]int, N+1)
  deleted := make([]int, N+1)


  for i := 0; i < N; i++ {
    fmt.Scan(&ancestor[i+1])
  }

  fmt.Scan(&C)
  tobeDeleted := make([]int, C)
  for i := 0; i < C; i++ {
    fmt.Scan(&t)
    //deleted[t] = 1
    tobeDeleted[i] = t
  }

  count := 0
  for i, v := range ancestor {
    fmt.Printf("%d--%d, ", i, v)
  }
  //fmt.Println("ancestor = ", ancestor)
  for i := C-1; i >= 0; i-- {
      k := tobeDeleted[i]
      if deleted[k] == 1 {
        continue;
      }
      fmt.Println("Deleting", k)
      for ; ancestor[k] != -1 && deleted[k] == 0; {
        fmt.Println("k =  ", k, "ancestor[k] = ", ancestor[k])
        deleted[k] = 1
        k = ancestor[k]
      }
      if ancestor[k] == -1 {
        fmt.Println("root")
        count++
      }
  }
  fmt.Println(count)
}
