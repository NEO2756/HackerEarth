/*
Gandhijee is interested in building human tree. He defined a human node as follows :

Person_Id = English alphabet {a...z} .
Person_Chain_Definition = Person_Id ( Person_Chain_Definition Person_Chain_Definition )
For example :
a( b(..) c( d( . e(..) ) f( . . ) ) ) refers to a human tree having a as the root human, whose chain links are are b and c, (spaces here are for clarity, input file has no spaces).

Note : After every Person_Id there is necessarily an opening parenthesis ( that starts the definition of the sub-tree under the person. For empty(NULL) human nodes, there is a '.' character.
Now, Gandhijee figure out the tree as follows :
Step 1 : He writes root human root id in column 0.
Step 2 : Then he writes left link's person id in 1 column to the left and right link's person id to 1 column to the right.
He continues writing the subtree, in a way that, for an id written in column C, it's left link's id is written in column C-1 and the right link's id is written in column C+1.

Now, Gandhijee wants to know the id's of people standing in a particular column C in lexicographically sorted order.
You are provided with column number C and the Tree definition string S. Gandhijee is very busy and hence he wants you to help him.

Input :
First line contains a number T, the number of test cases.
Each test case consists of a single line containing C, the column number for which you need to print the result, followed by a space followed by definition of the tree in the format given above, without any spaces.

Output :
For each test case, on a new line print the required id's in lexicographically sorted order, without spaces.
If there are no id's in column C, print "Common Gandhijee!" (quotes for clarity).

Constraints :
1 ≤ T ≤ 100
1 ≤ Len(S) ≤ 10000
-10000 ≤ C ≤ 10000

Image for given testcase.

*/

package main

import (
  "fmt"
  "sort"
)

type  node struct {
  name string
  val int
}
////a(c(f(.h(..))b(g(..).))e(.d(..)))
func main() {

  var T, C int
  var tree string
  var str, result []node
  var top node

  fmt.Scan(&T)
  for t := 0; t < T; t++ {

    fmt.Scan(&C)
    fmt.Scan(&tree)
    str = make([]node, 0)
    result = make([]node, 0)

    for _, item := range tree {
      char := string(item)

      if len(str) == 0 {
        str = append(str, node{char, 0})
        continue
      }

      top = str[len(str)-1]

      if char != ")" {
        if char == "(" {
          str = append(str, node{char, top.val})
          } else {
            if top.name == "(" {
              str = append(str, node{char, top.val - 1})  //left child
              } else {
                str = append(str, node{char, top.val + 2})  //right child
              }

            }

            } else {
              //fmt.Println("str = ", str)
              for i := 1; i < 4; i++ {
                e := str[len(str) - 1]
                if str[len(str)-1].name != string(".") && str[len(str)-1].name != string("(") {
                  result = append(result, e)
                }
                str = str[:len(str) - 1]
              }
              //fmt.Println("result", result)
              //fmt.Println()
            }
          }
          for len(str) > 0 {
            e := str[len(str) - 1]
            if str[len(str)-1].name != string(".") && str[len(str)-1].name != string("(") {
              result = append(result, e)
            }
            str = str[:len(str) - 1]
          }
          sort.SliceStable(result, func(i, j int) bool {
            return result[i].name < result[j].name
          })
          //fmt.Println("Final List=>", result)
          flag := false
          for _, v := range result {
            if v.val == C {
              flag = true
              fmt.Printf(v.name)
            }
          }
          if flag == false {
            fmt.Printf("Common Gandhijee!")
          }
          fmt.Println()
        }
      }
