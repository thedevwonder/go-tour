package basics

import "fmt"

func DeclareVarUninit() {
  var i int
  var f float64
  var b bool
  var s string

  fmt.Printf("%v %v %v %q\n", i, f, b, s)

}
