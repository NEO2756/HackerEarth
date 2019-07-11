/*


Mrinal is a coder from BVCOE. He always takes his laptop with him wherever he goes. One day (while on a trek trip ) a monster challenged him to solve one easy task else he was going to kill him.the task is as follows- Given a series of n numbers a1, a2, ..., an and a number of magic-queries. A magic-query is a pair (i, j) (1 ≤ i ≤ j ≤ n). For each magic-query (i, j), you have to return the number of distinct elements in the subsequence ai, ai+1, ..., aj.

Input

Line 1: n (1 ≤ n ≤ 30000). Line 2: n numbers a1, a2, ..., an (1 ≤ ai ≤ 10^6). Line 3: q (1 ≤ q ≤ 200000), the number of magic-queries. In the next q lines, each line contains 2 numbers i, j representing a magic-query (1 ≤ i ≤ j ≤ n).

Output

For each magic-query (i, j), print the number of distinct elements in the subsequence ai, ai+1, ..., aj in a single line. Example

SAMPLE INPUT
5
1 1 2 1 3
3
1 5
2 4
3 5
SAMPLE OUTPUT
3
2
3
*/

package main

import (
	"fmt"
	//"strings"
)
func main() {

  var N, Q, s, e int
  fmt.Scan(&N)

  arr := make([]int, N+1)
  for i := 1; i <= N; i++ {
    fmt.Scan(&arr[i])
  }

  fmt.Scan(&Q)
  for i := 0 ; i < Q; i++ {
    fmt.Scan(&s)
    fmt.Scan(&e)
    m := make(map[int]int)
    count := 0
    for ; s <= e; s++ {
       if _,ok := m[arr[s]]; !ok {
         m[arr[s]] = s
         //fmt.Println("s, e, map[s]", s, e, m[s], ok)
         count++
       }
    }
    fmt.Println(count)
  }
}
