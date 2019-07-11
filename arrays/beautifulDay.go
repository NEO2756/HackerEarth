/*The road has N checkpoints. Each checkpoint has a beauty value denoted as an array A in order from the beginning of the road to the end of it. The beauty
value of a checkpoint consists of the beauty of the places near the checkpoint. You are planning to t
ravel down the road. As it is a long road, you will stop at one of the checkpoints for some rest and after this rest, continue your journey.

So, your journey into two nearby sub-arrays,  being the checkpoints before your rest (including what you will stop) and  being the checkpoints after
 their rest, such that  (ie every checkpoint belongs to an exact sub-array).

The total beauty of the journey is defined by  at where It is the sum total of the beauty value in the sub-array S . You want to maximize the value of.

Input

The first entry line contains an integer , denoting the number of checkpoints. The second line contains N integers separated,
 denoting the beauty value of the checkpoints*/

 package main

 import (
 	"fmt"
 )

 func main() {

   var N int
   fmt.Scan(&N)

   arr := make([]int, N)

   for i := 0; i < N; i++ {
     fmt.Scan(&arr[i])
   }

   //sum1, sum2 := 0, 0

   sum1 := arr[0]
   sum2 := arr[N-1]
   //fmt.Println("sum1 = ", sum1, "sum2 = ", sum2)
   for i, j := 1, N-2; i <= j; {
     if sum1 < sum2 {
       sum1 += arr[i]
       i++
     } else {
       sum2 += arr[j]
       j--
     }
    // fmt.Println("sum1 = ", sum1, "sum2 = ", sum2, "i ,j", i, j)
   }
   fmt.Println(sum1*sum2)
 }
