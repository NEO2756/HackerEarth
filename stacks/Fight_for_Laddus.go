/*
Tuntun mausi challenged Bheem and his team to solve a problem. Raju, Chutki and Bheem are trying to solve this problem but are unable to do so. As you are a good friend of Raju, he asks for your help.

Given an array, For each element find the value of nearest element to the right which is having frequency greater than as that of current element. If there does not exist an answer for a position, then print '-1'

Please help Raju and his team to solve this problem to get the Laddus.

Input
First line contains T denoting the no of test cases.
First line of each test case contains N denoting the no of elements in array.
Second line of each test case contains N integers (A1..An) denoting the given array.

Output For each test case print space separated N numbers denoting the answer corresponding answer.

Constraints
20 points: T<=100 1<=N<=10^2 1<=Ai<=10^2

80 points: T<=100 1<=N<=10^5 1<=Ai<=10^5

Problem Setter
Aadersh Agarwal

Time Limit
2.5 seconds

SAMPLE INPUT
3
10
1 3 7 2 5 1 4 2 1 5
5
1 1 1 1 1
6
1 1 2 2 2 3
SAMPLE OUTPUT
-1 2 2 1 1 -1 2 1 -1 -1
-1 -1 -1 -1 -1
2 2 -1 -1 -1 -1

Explanation
Sample 1:
The given array is
1 3 7 2 5 1 4 2 1 5
3 1 1 2 2 3 1 2 3 2 denotes the frequency of ith element in array.
Index 1: there is no element to the right which has frequency greater than that of 1st element (i.e frequency >3)
Index 2: It has a frequency of 1. the nearest index to the right which has frequency greater than 1 is at index 4 i.e of element 2.

For last element the answer will always be -1
*/
package main

import (
  "fmt"
  "bufio"
  "os"
)
type node struct {
  v, a int
}

func toInt(buf []byte) (n int) {
    for _, v := range buf {
        n = n*10 + int(v-'0')
    }
    return
}
func print(s []*node) {
  for _, p := range s {
    fmt.Printf("%+v ", p)
 }
 fmt.Println()
}

func main() {
  var  T int
  fmt.Scan(&T)
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)

  for t := 0; t < T; t++ {
    scanner.Scan()
    n := toInt(scanner.Bytes())

    arr := make([]node, n)
    frequency := make([]int, 1000001)
    // answer := make([]int, n)
    stack := make([]*node, 0)

    for i := 0; i < n; i++ {
      scanner.Scan()
      val := toInt(scanner.Bytes())
      arr[i].v = val
      //fmt.Println(val)
      frequency[val] += 1
    }

    for i := 0; i < n; {

        if len(stack) == 0 {
          stack = append(stack, &arr[i])
          i++
          continue
        }

        top := (stack[len(stack) - 1]).v

        //fmt.Println("top = ", top, "new = ", arr[i].v)
        //print(stack)
        if frequency[top] >= frequency[arr[i].v] {
          stack = append(stack, &arr[i])
          i++
        } else {
          stack[len(stack) - 1].a = arr[i].v
          stack = stack[:len(stack) - 1]
        }
        //fmt.Printf("After stack = ")
        //print(stack)
    }
    for _, nod := range arr {
      if nod.a == 0 {
        fmt.Printf("-1 ")
      } else {
        fmt.Printf("%d ", nod.a)
      }
    }
    frequency = nil
    fmt.Println()
  }
}
