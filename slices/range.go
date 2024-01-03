package slices

import "fmt"

var s = []int {1, 2, 4, 8, 16, 32, 64}

func UseRange() {
  for i, v := range s {
    fmt.Printf("2**%d = %d\n", i, v)
  }
}

func UseShortRange() {
  //Print values
  for _, v := range s {
    fmt.Println(v)
  }

  //Print index
  for i := range s {
    fmt.Println(i)
  }
}
