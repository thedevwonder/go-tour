package slices

import "fmt"

func CreateSliceLiteral() {
  s := []struct {
        X int
        Y bool
      }{
	  {1, true},
	  {2, true},
	  {3, false},
	  {4, false},
       }
  fmt.Println(s)
}
