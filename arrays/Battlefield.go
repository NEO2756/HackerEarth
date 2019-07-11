/*
The war goes on.The Knights and the Demons are fighting at their fullest but as expected, the evil is dominating the good ones.In order to defeat the Demons the knights have to come together.
Initially everyone is standing in a circular path at the top of the Demon tower.The warrior at index 1 is standing next to warrior at index n and before the warrior at index 2.In order to win the fight the knights have to come together.
The knights have a special power through which they can swap them self with anyone.
Help the Knights decide the minimum number of swaps they have to do so that all of them stand together.

Input
The first line of the input contains t (the number of test cases).
The next 2*t line contains the following:
The first line contains n (the size of string).
The next line contains the string s.
The ith character denotes if there is a Knight (K) or a Demon (D).

Output
The output contains an integer denoting the minimum number of swaps required for each test case t.


Constraints
1<=t<=100
1<=n<=105

*/


package main

import (
	"fmt"
//  "os"
//  "bufio"
	//"strings"
)

  func main() {

  var T, N, ones, max int

  fmt.Scanf("%d", &T)
  for t := 0; t < T; t++ {
    max = -1
    ones = 0
    var text string
    fmt.Scanf("%d", &N)
    arr := make([]int, N)
    fmt.Scanf("%s", &text)
    for i, c := range text {
        if string(c) ==  "D" {
          arr[i] = 0
        } else if string(c) == "K"{
          arr[i] = 1
          ones++
        }
    }
      ws := ones;
      sum := 0
      //first window
      for k := 0; k < ws; k++ {
        sum += arr[k]
      }
      //start from secound window
      max = sum
      for k := 1; k < N; k++ {
        sum = sum - arr[(k - 1)%N] + arr[(k + ws - 1)%N]
        //fmt.Println("k = ", k%N, "k + ws - 1 = ", (k + ws - 1)%N, "sum = ", sum)
        if sum > max {
          max = sum
        }
       }
       fmt.Println(ones - max)
      }

  }
