package slices

import "fmt"

func Show_Nil_Slices() {
  var s []int
  fmt.Println(len(s), cap(s))
}
