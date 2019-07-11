/*
It's sweet distribution day in Micro's school. He's very happy. All the students in Micro's class are sitting on chairs which are arranged in a matrix of size  i.e. there are N rows of chairs numbered from 1 to N and in each row there are M chairs numbered from 1 to M. Micro is sitting at coordinate  ( chair of  row). Teacher gives the box to a student sitting in one of the four corners: , ,  or . Students have to take one sweet from the box and pass the box to the next student (student sitting to left, right, front or back). For a student sitting at coordinate , he'll follow the following priority order:

If there is a student in the Right who has not received sweet, then pass it right .
If there is a student in the Left who has not received sweet, then pass it left .
If there is a student in the Front who has not received sweet, then pass it front .
If there is a student in the Back who has not received sweet, then pass it back .
Shout Over, meaning that all students have received sweets.
Now, Micro is curious to find out the direction in which he'll have to pass the sweet box. Since there are a lot of students in Micro's class, it will take long for the box to reach him, and you know Micro, he just can't wait. So he asks you to find out the direction in which he'll have to pass the box, or will he have to shout Over.

Input:
First line consists of a single integer, T, denoting the number of test cases. Each test case consists of three lines.
First line of each test consists of two space separated integer denoting N and M.
Second line of each test case consists of two space separated integers  and , denoting the coordinate of the corner from which sweet distribution starts.
Third line of each test case consists of two space separated integers denoting  and , coordinates of Micro.

Output:
For each test case, print "Left", "Right", "Front", "Back", "Over", depending on whether Micro will pass the box to student in left, right, front or back, or if he'll shout over.*/


package main

import (
	"fmt"
)

func main() {
  //var visited [1000][1000]int8
	var T, N, M, s1, s2, mx, my int
  fmt.Scan(&T)
  for t := 0; t < T; t++ {
    fmt.Scan(&N)
    fmt.Scan(&M)
    fmt.Scan(&s1)
    fmt.Scan(&s2)
    fmt.Scan(&mx)
    fmt.Scan(&my)
    //fmt.Println("N, M, s1, s2, mx, my", N, M, s1, s2, mx, my)

		if s1 == 1 && s2 == 1 {
			if mx == N && my == M && N%2 != 0 {
				fmt.Println("Over")
				continue
			}

			if mx == N && my == 1 && N%2 == 0 {
				fmt.Println("Over")
				continue
			}

			if my == 1 {   //first column
				if mx == N && N%2 != 0 {  // 0, N-1
					fmt.Println("Right")
					continue
				} else {
					fmt.Println("Front")
					continue
				}
			} else if my == M {  //0 , M-1
				if mx%2 != 0 {
					fmt.Println("Front")
					continue
				} else {
					fmt.Println("Left")
					continue
				}
			} else {  //some where in between, so depends upon row
				if mx%2 != 0 {
					fmt.Println("Right")
					continue
				} else {
					fmt.Println("Left")
					continue
				}
			}
		}


		if s1 == 1 && s2 == M {

			if N%2 == 0 {
				if mx == N && my == 1 {
					fmt.Println("Over")
					continue
				}
				if mx == N && my == M {
					fmt.Println("Left")
					continue
				}
			} else {
				if mx == N && my == 1 {
					fmt.Println("Over")
					continue
				}
				if mx == N && my == M {
					fmt.Println("Left")
					continue
				}
			}
			if mx == N && my == 1 && N%2 == 0 {
				fmt.Println("Over")
				continue
			}
			if mx == N && my == M && N%2 != 0 {
				fmt.Println("Over")
				continue
			}

			if my == M {   //first column
				if mx == N {  // 0, N-1
					fmt.Println("Left")
					continue
				} else {
					fmt.Println("Front")
					continue
				}
			} else if my == 1 {  //0 , M-1
				if mx%2 != 0 {
					fmt.Println("Front")
					continue
				} else {
					fmt.Println("Right")
					continue
				}
			} else {  //some where in between, so depends upon row
				if mx%2 != 0 {
					fmt.Println("Left")
					continue
				} else {
					fmt.Println("Right")
					continue
				}
			}
		}

		if s1 == N && s2 == 1 {
			if mx == 1 && my == M && N%2 != 0 {
				fmt.Println("Over")
				continue
			}
			if mx == 1 && my == 1 && N%2 == 0 {
				fmt.Println("Over")
				continue
			}

			if my == 1 {   //first column
				if mx == 1 {  // 0, N-1
					fmt.Println("Right")
					continue
				} else {
					fmt.Println("Back")
					continue
				}
			} else if my == M {  //0 , M-1
				if mx%2 != 0 {
					fmt.Println("Back")
					continue
				} else {
					fmt.Println("Left")
					continue
				}
			} else {  //some where in between, so depends upon row
				if mx%2 != 0 {
					fmt.Println("Right")
					continue
				} else {
					fmt.Println("Left")
					continue
				}
			}
		}

		if s1 == N && s2 == M {
			if mx == 1 && my == 1 && N%2 != 0{
				fmt.Println("Over")
				continue
			}
			if mx == 1 && my == M && N%2 == 0{
				fmt.Println("Over")
				continue
			}

			if my == M {   //first column
				if mx == 1 {  // 0, N-1
					fmt.Println("Left")
					continue
				} else {
					fmt.Println("Back")
					continue
				}
			} else if my == 1 {  //0 , M-1
				if mx%2 != 0 {
					fmt.Println("Back")
					continue
				} else {
					fmt.Println("Right")
					continue
				}
			} else {  //some where in between, so depends upon row
				if mx%2 != 0 {
					fmt.Println("Left")
					continue
				} else {
					fmt.Println("Right")
					continue
				}
			}
		}

  }
}
