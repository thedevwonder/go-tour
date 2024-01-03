package slices

import "fmt"

func GetSliceDefaults() {
  s := []int {1,2,3,4,5}
  
  s = s[:4]
  fmt.Println(s)

  s = s[1:]
  fmt.Println(s)
}
