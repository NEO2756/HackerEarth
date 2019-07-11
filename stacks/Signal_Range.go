/*
Various signal towers are present in a city.Towers are aligned in a straight horizontal line(from left to right) and each tower transmits a signal in the right to left direction.Tower A shall block the signal of Tower B if Tower A is present to the left of Tower B and Tower A is taller than Tower B. So,the range of a signal of a given tower can be defined as :

{(the number of contiguous towers just to the left of the given tower whose height is less than or equal to the height of the given tower) + 1}.

You need to find the range of each tower.

INPUT

First line contains an integer T specifying the number of test cases.

Second line contains an integer n specifying the number of towers.

Third line contains n space separated integers(H[i]) denoting the height of each tower.

OUTPUT

Print the range of each tower (separated by a space).

Constraints

1 <= T <= 10

2 <= n <= 10^6

1 <= H[i] <= 10^8

SAMPLE INPUT
1
7
100 80 60 70 60 75 85
SAMPLE OUTPUT
1 1 1 2 1 4 6
Explanation
6th tower has a range 4 because there are 3 contiguous towers(3rd 4th and 5th towers) just to the left of the 6th tower whose height is less than height of 6th tower.*/


package main

import (
  "fmt"
  "bufio"
  "os"
)

type node struct {
  v, r int
}

func toInt(buf []byte) (n int) {
    for _, v := range buf {
        n = n*10 + int(v-'0')
    }
    return
}

func main() {
  var T, N int

  fmt.Scan(&T)
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)
  for t := 0; t < T; t++ {

    scanner.Scan()
    N = toInt(scanner.Bytes())
    arr := make([]node, N)
    stack := make([]node, 0)

    for i := 0; i < N; i++ {
      scanner.Scan()
      val := toInt(scanner.Bytes())
      arr[i].v = val
    }

    for i := 0; i < len(arr); {
      if len(stack) == 0 {
        arr[i].r += 1
        stack = append(stack, arr[i])
        i++
        continue
      }

      top := stack[len(stack) - 1]
      //fmt.Println("before =", stack, "arr[i] = ", arr[i])
      if top.v > arr[i].v {
        arr[i].r += 1
        stack = append(stack, arr[i])
        i++
      } else {   // <=
        e := stack[len(stack) - 1]
        stack = stack[:len(stack)-1] //pop
        arr[i].r +=  e.r
      }
      //fmt.Println("after =", stack)
    }
    for _, v := range arr {
      fmt.Printf("%d ", v.r)
    }
    fmt.Println()
    //fmt.Println("final =", arr)
  }
}
