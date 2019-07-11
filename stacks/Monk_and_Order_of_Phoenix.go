package main

import (
  "fmt"
  "bufio"
  "os"
)

var arr  [][]int

func toInt(buf []byte) (n int) {
    for _, v := range buf {
        n = n*10 + int(v-'0')
    }
    return
}

type queryType struct {
	typ int64
	k   int64
	h   int64
}

func calculateBound (target, i int ) (int) {
//  for i := 1; i < N; i++ {
     low := 0
     high := len(arr[i]) - 1
     for low != high {
         mid := (low + high) / 2
         if arr[i][mid] <= target {
           //fmt.Println("here", arr[i][mid], target)
           low = mid + 1
        } else {
          high = mid
        }
     }
     //fmt.Println("low = ", low, "high =" , high, arr[i][low], target)
  //}
  if arr[i][low] > target {
    return arr[i][low]
  } else {
    return -1
  }
}

func main() {
  var N int
  var flag bool

  fmt.Scan(&N)
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)
  arr = make([][]int, N)
  for i := 0; i < N; i++ {
    scanner.Scan()
    n := toInt(scanner.Bytes())
    arr[i] = make([]int, n)
    for j := 0;  j < n;  j++ {
      scanner.Scan()
      arr[i][j] = toInt(scanner.Bytes())
    }
  }
  //fmt.Println(arr)
  scanner.Scan()
  Q := toInt(scanner.Bytes())

  queries := make([]queryType, Q)
	for i = 0; i < Q; i++ {
		scanner.Scan()
		qVal, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		switch qVal {
		case 0:
			scanner.Scan()
			k, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			queries[i] = queryType{typ: 0, k: k, h: 0}
		case 1:
			scanner.Scan()
			k, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			scanner.Scan()
			h, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			queries[i] = queryType{typ: 1, k: k, h: h}
		case 2:
			queries[i] = queryType{typ: 2, k: 0, h: 0}
		}
	}

  for _, x := range queries {
    scanner.Scan()
    x := toInt(scanner.Bytes())

    if x == 0 {
      scanner.Scan()
      k := toInt(scanner.Bytes())
      k = k - 1
      arr[k] = arr[k][:len(arr[k]) - 1]
    } else if x == 1 {
      scanner.Scan()
      k := toInt(scanner.Bytes())
      scanner.Scan()
      h := toInt(scanner.Bytes())
      k = k - 1
      arr[k] = append(arr[k], h)
      } else { // 2
        //fmt.Println(arr)
        for _, v := range arr[0] {
          flag = true
          for e := 1; e < N; e++ {
            v = calculateBound(v, e)
            //fmt.Println("v = ", v)
            if v == -1 {
              flag = false
              break
            }
          }
          if flag == true {
            break
          }
        }
        if flag == true {
          fmt.Println("YES")
        } else {
          fmt.Println("NO")
        }
      }

    }
}
