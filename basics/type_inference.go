package basics

import "fmt"

func InferType() {
  var i int
  //type is infered from right hand side of the equation
  j := i
  fmt.Printf("Type of j is %T ", j)
}

