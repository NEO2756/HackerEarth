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
  //fmt.Println(this.length)
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

/////MAIN code
func main() {

  var T, N int
  fmt.Scanf("%d", &T)

  for t := 0; t < T; t++ {
    fmt.Scan(&N)
    arr := make([]int, N+1)
    front := make([]int, N+1)
    back := make([]int, N+1)

    for i := 1; i <= N; i++ {
      fmt.Scan(&arr[i])
    }

    max, ans, top := 0, 0, 0

    s := New()

    s.Push(1)
    front[N], back[1] = N, 1 //back of 1st driver and front of last driver is always zero
    for i := 2; i <= N; {
      if s.Len() > 0 {
        top = s.Peek().(int)
      }
      if s.Len() == 0 || arr[i] <= arr[top] {
        s.Push(i)
        //back[i] = top
        front[top] = i
        i++
      } else {
        e := s.Pop().(int)
        fmt.Println("pop e = ", arr[e])
        back[i] = back[e]
        front[e] = i
        tmp := (e) * (((front[e] - e) + (e - back[e]))%1000000007)
      fmt.Println("here", tmp, max, "front = ", (front[e] - e), "back = ", (e - back[e]))
        if tmp > max {
          max = tmp
          ans = e
        }
      }
    }
      fmt.Println("front =", front, "back = ", back)
      min_idx := s.Peek().(int)
      for s.Len() != 0 {
        e := s.Pop().(int)
        fmt.Println("poped = ", arr[e])
        if s.Len() > 0  {
          if arr[e] == arr[s.Peek().(int)] {
            front[s.Peek().(int)] = e
          } else {
            front[s.Peek().(int)] = min_idx
          }
        }

        tmp := (e) * (((front[e] - e) + (e - back[e])%1000000007))
        fmt.Println(tmp, max, "front = ", (front[e] - e), "back = ", (e - back[e]))
        if tmp > max {
          max = tmp
          ans = e
        }
    }
     fmt.Println("test case ", t , ans)
  }
}
