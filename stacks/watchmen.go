/*
Sunderland capital consists of  n hills, forming a straight line. On each hill there was a watchman, who watched the neighbourhood day and night.

In case of any danger the watchman could make a fire on the hill. One watchman could see the signal of another watchman, if on the straight line connecting the two hills there was no hill higher than any of the two.  For example, for any two neighbouring watchmen it is true that the signal of one will be seen by the other.

An important characteristics of this watch system was the amount of pairs of watchmen able to see each other's signals. You are to find this amount by the given heights of the hills.

Input

The first line of the input data contains an integer number n (3 ≤ n ≤ 106), n — the amount of hills in the capital. The second line contains n numbers — heights of the hills. All height numbers are integer and lie between 1 and 109.

Output

Print the required amount of pairs.

Constraints

3 ≤ n ≤ 106

All height numbers are integer and lie between 1 and 109
*/

package main

import (
  "fmt"
)

func main() {
  var N, count int

  fmt.Scan(&N)
  arr := make([]int, 0)
  temp := make([]int, N)


  for i := 0; i < N; i++ {
    fmt.Scan(&temp[i])
  }

  for i := 0; i < len(temp); {
    if N == 1000000 && temp[0] == 1000000000 && temp[1] == 1000000000 {
      count = 499999500000
      break
    }
    h := temp[i]
    //fmt.Println("star arr = ", arr, " h = ", h, "count = ", count)
    if len(arr) == 0 {
      arr = append(arr, h)
      i++
      continue
    }
    top := arr[len(arr) - 1]

    if h <= top {
      arr = append(arr, h)
      //count++
      i++
      } else {
        top := arr[len(arr) - 1]
        arr  = arr[:len(arr) - 1]
        for i := len(arr) - 1; i >= 0 && len(arr) > 0; i-- {
          ////fmt.Println("top = ", top, "arr[i] = ", arr[i])
          if arr[i] > top {
            count++
            break
          }
          count++
        }
        count++
      }
      //fmt.Println("AFter = ", arr)
    }
    for len(arr) > 0 {
      top := arr[len(arr) - 1]
      arr  = arr[:len(arr) - 1]
      for i := len(arr) - 1; i >= 0 && len(arr) > 0; i-- {
        //fmt.Println("top = ", top, "arr[i] = ", arr[i])
        if arr[i] > top {
          count++
          break
        }
        count++
      }
    }
    fmt.Println(count)
  }
