/**
There are N frustrated coders standing in a circle with a gun in their hands. Each coder has a skill value S[ i ] and he can only kill those coders that have strictly less skill than him. One more thing, all the guns have only 1 bullet. This roulette can take place in any random order. Fortunately, you have the time stone (haaan wo harre wala) and you can see all possible outcomes of this scenario. Find the outcome where the total sum of the remaining coder's skill is minimum. Print this sum.

INPUT

The first line contains N the no. of coders

The next line contains N  elements where the ith element is theS[ i ] of ith coder.

OUTPUT

Print a single line containing the minimum sum

CONSTRAINTS

1<= N <= 1000000

1<=S[ i ]<=1000
*/

package main

import (
  "fmt"
  "sort"
  "bufio"
  "os"
  "strings"
  "strconv"
)

func main() {
  var N int

  fmt.Scan(&N)
  arr := make([]int, N)
  nums, _ := bufio.NewReader(os.Stdin).ReadString('\n')
  nums = strings.TrimSuffix(nums, "\n")
  ns := strings.Split(nums, " ")
  for _, i := range ns {
    r, _ := strconv.Atoi(i)
    arr = append(arr, r)
  }

  // for i := 0; i < N; i++ {
  //   fmt.Scan(&arr[i])
  // }

  sort.Ints(arr)
  //fmt.Println(arr)

  for k , v := range arr {
    for i := k; i >= 0; i-- {
      if arr[i] < v && arr[i] != 0 {
        arr[i] = 0
        break
      }
    }
  }
  //fmt.Println(arr)
  sum := 0
  for _, v := range arr {
    sum += v
  }

  fmt.Println(sum)
}
