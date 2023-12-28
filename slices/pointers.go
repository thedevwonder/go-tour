package slices

import "fmt"

func DeclarePointers() {
  i, j := 42, 2701
  
  p := &i
  fmt.Printf("printing i:%d & j:%d \n", i, j)

  *p = 21
  fmt.Println("set i using pointer", i)

  p = &j
  *p = *p / 37

  fmt.Println("set j using pointer", j)
}
