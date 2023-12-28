package slices

import "fmt"

var (
  v1 = NewStruct{1, 2}
  v2 = NewStruct{X: 3}
  v3 = NewStruct{}
  p = &NewStruct{4, 5}
)

func DeclareStructLit() {
    fmt.Println(v1, v2, v3, p)
}
  
