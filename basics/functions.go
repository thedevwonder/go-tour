package basics

import "fmt"

func add(x int, y int) int {
 /* type is written after 
    the variable */
  return x + y
}

func CallAdd(x, y int) {
  fmt.Println("type is shared in the func argument", add(x, y))
}
