package slices

import "fmt"

func GetLenCap() {
  s := []int {1, 2, 4, 5, 6}

  s = s[:0]
  fmt.Println(s)

  s = s[:4]
  fmt.Println(s)

  fmt.Printf("len = %d, cap = %d", len(s), cap(s))
}
