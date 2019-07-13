package main

import (
  "fmt"
  "bufio"
  "os"
  //"sort"
  //"strconv"
)

func toInt(buf []byte) (n int) {
    for _, v := range buf {
        n = n*10 + int(v-'0')
    }
    return
}

func calculateBound (target int, arr []int ) (int) {

     ans := -1
     low := 0
     high := len(arr) - 1
     for low <= high {
         mid := (low + high) / 2

         if arr[mid] >= target {
           high = mid - 1
        } else {
          ans = mid
          low = mid + 1
        }
     }
     //fmt.Println("LOW = ", low, "HIGH =" , high, ans, target)
     if ans != -1 {
       return arr[ans]
     }
     return ans
}

func main() {
  var N int
  fmt.Scan(&N)
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)

  arr := make([]int, N)
  stack := make([]int, 0)
  answer := make([]int, N)
  for i := 0 ; i < N; i++ {
    scanner.Scan()
    val := toInt(scanner.Bytes())
    arr[i] = val
  }

  for i := 0 ; i < N; {

      if len(stack) == 0 {
        stack = append(stack, arr[i])
        answer[i] = -1
        i++
        continue
      }

      top := stack[len(stack) - 1]
      //fmt.Println("stack = ", stack, "arr[i] =", arr[i], "top = ", top)
      if top > arr[i] {
        stack = stack[:len(stack) - 1]
      } else {
        v := calculateBound(arr[i], stack)
        answer[i] = v
        stack = append(stack, arr[i])
        i++
      }
     //fmt.Println("After stack = ", stack)
  }
  for _, v := range answer {
    fmt.Printf("%d ", v)
  }
}
