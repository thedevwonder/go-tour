package slices

import "fmt"

type NewStruct struct {
  X int
  Y int
}

func DeclareStruct() {
  fmt.Println(NewStruct{1, 2})
}

func GetX() {
  v := NewStruct{1, 2}
  fmt.Println(v.X)
}
