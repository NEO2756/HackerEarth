/*
It's sweet distribution day in Micro's school. He's very happy. All the students in Micro's class are sitting on chairs which are arranged in a matrix of size  i.e. there are N rows of chairs numbered from 1 to N and in each row there are M chairs numbered from 1 to M. Micro is sitting at coordinate  ( chair of  row). Teacher gives the box to a student sitting in one of the four corners: , ,  or . Students have to take one sweet from the box and pass the box to the next student (student sitting to left, right, front or back). For a student sitting at coordinate , he'll follow the following priority order:

If there is a student in the Right who has not received sweet, then pass it right .
If there is a student in the Left who has not received sweet, then pass it left .
If there is a student in the Front who has not received sweet, then pass it front .
If there is a student in the Back who has not received sweet, then pass it back .
Shout Over, meaning that all students have received sweets.
Now, Micro is curious to find out the direction in which he'll have to pass the sweet box. Since there are a lot of students in Micro's class, it will take long for the box to reach him, and you know Micro, he just can't wait. So he asks you to find out the direction in which he'll have to pass the box, or will he have to shout Over.

Input:
First line consists of a single integer, T, denoting the number of test cases. Each test case consists of three lines.
First line of each test consists of two space separated integer denoting N and M.
Second line of each test case consists of two space separated integers  and , denoting the coordinate of the corner from which sweet distribution starts.
Third line of each test case consists of two space separated integers denoting  and , coordinates of Micro.

Output:
For each test case, print "Left", "Right", "Front", "Back", "Over", depending on whether Micro will pass the box to student in left, right, front or back, or if he'll shout over.*/


package main

import (
	"fmt"
)


var T, N, M, s1, s2, mx, my int

var idx = [4]int{0, 0, -1, 1}
var idy = [4]int{1, -1, 0, 0}

func printAnswer(visited [][]int8) {
  for i := 0; i < 4; i++ {
     newX := mx + idx[i]
     newY := my + idy[i]
     if newX >= 0 && newX < N && newY >=0 && newY < M && visited[newX][newY] != 1 {
       if i == 0 {
         fmt.Println("Right")
       } else if (i == 1) {
         fmt.Println("Left")
       } else if (i == 2) {
         fmt.Println("Front")
       } else {
         fmt.Println("Back")
       }
       return
     }
   }
   fmt.Println("Over")
}

func dfs(visited [][]int8, x, y int) {

    if x == mx && y == my {
      printAnswer(visited);
      return;
    }
    //fmt.Println(x, y)
    for i := 0; i < 4; i++ {
       newX := x + idx[i]
       newY := y + idy[i]
       if newX >= 0 && newX < N && newY >=0 && newY < M && visited[newX][newY] != 1 {
         visited[newX][newY] = 1
         dfs(visited, newX, newY)
       }
    }
}



func main() {
  //var visited [1000][1000]int8
  fmt.Scan(&T)
  for t := 0; t < T; t++ {
    fmt.Scan(&N)
    fmt.Scan(&M)

    visited := make([][]int8, N)
    for i := range visited {
      visited[i] = make([]int8, M)
    }

    fmt.Scan(&s1)
    fmt.Scan(&s2)
    fmt.Scan(&mx)
    fmt.Scan(&my)
    s1 -= 1
    s2 -= 1
    mx -= 1
    my -= 1
    //fmt.Println(N, M, "start from = ", s1, s2, "micro = ",  mx, my)
    visited[s1][s2] = 1
    dfs(visited, s1, s2)

  }


}
