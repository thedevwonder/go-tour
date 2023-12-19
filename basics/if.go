package basics

import "fmt"

func Condition(x int) {
  if x < 0 {
    fmt.Println("Negative")
    return
  }
  fmt.Println("Non Negative")
  return 
}

func ShortCondition (x int) {
  if y := 5; x + y < 0 {
    fmt.Println("Negative")
    return
  }
  fmt.Println("Non Negative")
}  
