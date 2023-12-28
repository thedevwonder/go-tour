package slices

import "fmt"

func DeclareStructP() {
  v := NewStruct{1, 2}
  p := &v
  p.X = 1e9
  fmt.Println(v)
}
