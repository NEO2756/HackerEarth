package main

import (
	"fmt"
	//"strings"
)

func seattype(dis int) (string) {
	if dis == 0 || dis == 5 || dis == 6 || dis == 11 {
		return "WS"
	} else if dis == 1 || dis == 4 || dis == 7 || dis == 10 {
		 return "MS"
	} else if dis == 2 || dis == 3 || dis == 8 || dis == 9 {
		 return "AS"
	}
	return "error"
}

func main() {

  var T, N, coach_number int

  fmt.Scanf("%d", &T)
  for i := 0; i < T; i++ {
    fmt.Scanf("%d", &N)
    if N%12 == 0 {
			coach_number = N/12 - 1
		} else {
			coach_number = N/12
		}
    start_seat := 12 * coach_number + 1
		//fmt.Println(start_seat)
    //next_seat := start_seat
    fmt.Printf("%d %s\n", (start_seat + 11) - (N - start_seat), seattype(N-start_seat))
    }
}
