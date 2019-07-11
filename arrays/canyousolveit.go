package main

import (
	"fmt"
	"math"
)

func main() {

	var T, n, temp int
	fmt.Scanf("%d", &T)
	for i := 0; i < T; i++ {
		fmt.Scanf("%d", &n)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &arr[j])
		}

    //fmt.Println(arr, n)
    sum := int(math.Abs(float64(arr[0] - arr[n-1])) + float64(n-1))
    //fmt.Println("sum == ", sum)
    for start , end := 0, n-1; start < end; {
    // arr[start] <=  arr[end-1] >= arr[end], continue, it wont give us sum
    if ((arr[end - 1] <= arr[end] && arr[end - 1] >= arr[start]) || (arr[end - 1] >= arr[end] && arr[end - 1] <= arr[start])) {
    //   end--
       //continue
    } else {
    //  end--
      temp = int(math.Abs(float64(arr[start] - arr[end - 1])) + math.Abs(float64(start - (end - 1 ))))
      fmt.Println("1- start = ", start, "end = ", end - 1, "sum = ", sum, "temp = ", temp)
      if (temp > sum) {
        sum = temp
      }
    }
    // arr[start] <= arr[start+1] >= arr[end], continue, it wont give us sum
    if ((arr[start+1] >= arr[start] && arr[start + 1] <= arr[end]) || (arr[start + 1] <= arr[start] && arr[start+1] >= arr[end])) {
      //start++
      //continue
    } else {
    //  start++
      temp = int(math.Abs(float64(arr[start + 1] - arr[end])) + math.Abs(float64(start + 1 - end)))
      fmt.Println("2- start = ", start + 1, "end = ", end, "sum = ", sum, "temp = ", temp)
      if (temp > sum) {
        sum = temp
      }
    }
    start++
    end--
    // fmt.Println("temp = ", temp, "sum = ", sum)
    // if temp > sum {
    //   sum = temp
    // }
	}
  fmt.Println(sum)
 }
}
