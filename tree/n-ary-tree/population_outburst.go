/*
A new species is trying to rule the planet. This species is creating their own population outburst to dominate other species. It all started with 1 single member of the species. The population increases in treelike fashion abiding by few rules as listed below.

Single member is able to reproduce by itself.
A new member is added to the population every minute.
Every member is associated with integral name.
Multiple members can share a common name.
Every member has it's own reproduction capacity, that is maximum number of children it can reproduce.
A member can start to reproduce only if all members older than it have exhausted their reproduction capacity.
Level 0 in family tree of this species comprise of single member at the start of multiplication.
Integral name of single member at the start is 0.
The population grows level wise, where number of members at level i is dependent on reproduction capacity of members at prior level.
Given the integral name of new member and it's reproduction capacity that is added to the population, you have to find it's parent, level at which it is added and it's ascending age wise rank among siblings.

Input:
First line of the input contains 2 integers, , representing number of minutes we will be examining the population increase and reproduction capacity of member at epoch. Next N line contains 2 integers each, , representing integral name and reproduction capacity of new member born at time i.

Output:
N lines, each line containing 3 integers, , representing integral name of the parent, level at which it is added and it's ascending age wise rank among siblings.

Note :
It will always be possible to reproduce a new child or in other words, through out the given time, there exists atleast one member which can still accomodate new child.
*/


package main

import (
  "fmt"
  //"math"
)
type node struct {
  name, capacity int
}

func main() {

  var N, capacity0, name, capacity int

  fmt.Scan(&N)
  fmt.Scan(&capacity0)
  var q2 []node
  q1 := make([]node, 0)
  q1 = append(q1, node{0, capacity0})

  level := 0

  for N > 0 {
    level++
    //fmt.Println("capacity = ",  e.capacity)
    for len(q1) != 0 && N > 0 {
      e := q1[0] //dqueue
      q1 = q1[1:]
      for rank := 1; rank <= e.capacity && N > 0; rank++ {
        fmt.Println(e.name , level, rank)
        fmt.Scan(&name)
        fmt.Scan(&capacity)
        q2 = append(q2, node{name, capacity})
        N--
      }
   }
   q1 = append(q1[0:], q2...)
   q2 = nil
   //fmt.Println(q1)
 }


 //  for N > 0 {
 //    //fmt.Println(len(q1))
 //    level++
 //    e := q1[0] //dqueue
 //    q1 = q1[1:]
 //    fmt.Println("capacity = ",  e.capacity)
 //    for e.capacity > 0 {
 //      rank := 1
 //      fmt.Println(e.name, "", level, "", rank)
 //      fmt.Scan(&name)
 //      fmt.Scan(&capacity)
 //      q2 = append(q2, node{name, capacity})
 //      rank++
 //      N--
 //      e.capacity--
 //   }
 //   q1 = append(q1[0:], q2...)
 //   q2 = nil
 //   fmt.Println(q1)
 // }


}
