/*
There are N warehouses. The warehouses are located in a straight line and are indexed from 1 to N. Each warehouse contains some number of sacks.

A thief decides to rob these warehouses. Thief figured out that he can escape the police if and only if he follows both the following 2 constraints:

He will rob only one continuous segment of warehouses.
He will rob same number of sacks from each warehouse.
Thief wants to calculate the maximum number of sacks he can steal without getting caught by the police.

Input Format:

The first line contains an integer T denoting number test cases.

The first line of each test case contains a single integer N denoting number of warehouses.

The second line of each test case contains N space-separated integers : denotes number of sacks in warehouse.

Constraints:

Output Format:

Output exactly T lines.

The  line should contain a single integer, i.e the answer for  testcase(maximum number of sacks thief can steal without getting caught).
*/


package main

import (
	"fmt"
	//"github.com/golang-collections/collections/stack"
)
type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}
)
// Create a new stack
func New() *Stack {
	return &Stack{nil,0}
}
// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}
// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}
// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}
// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value,this.top}
	this.top = n
	this.length++
}
func main() {

  var T, N, sum int
  fmt.Scanf("%d", &T)

  for t := 0; t <  T; t++ {
    fmt.Scan(&N)
    sum = 0
    arr := make ([]int, N)
    for i := 0; i < N; i++ {
      fmt. Scan(&arr[i])
    }

    s := New()
    i := 0

    for i = 0; i < N;  {
      if s.Len() == 0 || arr[s.Peek().(int)] <= arr[i] {
				if (s.Len() != 0) {
					//fmt.Println("Top", arr[s.Peek().(int)], "<=" , "arr[i]", arr[i]  )
				}
				//fmt.Println("pushing: ", i)
        s.Push(i)
				i++
      } else {
        e := s.Pop().(int)
				//fmt.Println("poped: ", e)
        n := 0
				if s.Len() == 0 {
					n = i
				} else {
					n = i - s.Peek().(int) - 1
				}
				tmp := arr[e] * n

				if tmp > sum {
					sum = tmp
				}
				//fmt.Println("1 - sum = ", sum, "temp =", tmp)
      }
    }

		for s.Len() != 0 {
			e := s.Pop().(int)
			n := 0
			if s.Len() == 0 {
				n = i
			} else {
				n = i - s.Peek().(int) - 1
			}
			tmp := arr[e] * n

			if tmp > sum {
				sum = tmp
			}
			//fmt.Println("1 - sum = ", sum, "temp =", tmp)
		}
    fmt.Println(sum)
  }
}
