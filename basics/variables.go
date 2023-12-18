package basics

import "fmt"

var c, python, java bool

func DeclareVariable () {
  var i int
  fmt.Println(i, c, python, java)
}

func InitializeVariables () {
  var c, python, java = true, false, "no!"
  fmt.Println(c, python, java)
}

func ShortVariable () {
  k := 3
  fmt.Println(k)
}
