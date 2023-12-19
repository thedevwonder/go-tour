package basics

import (
  "fmt"
  "runtime"
)

func GetOS() {
  fmt.Println("Go is running on: ")
  switch os:= runtime.GOOS; os {
  case "darwin":
    fmt.Println("OSX")
  case "linux":
    fmt.Println("Linux")
  default: 
    fmt.Printf("%s \n", os)
  }
}


/* evaluation order
   is top to bottom */ 
func GetNumber(x int) {
  switch x {
  case 1:
    fmt.Println("One")
  case 2:
    fmt.Println("Two")
  default: 
    fmt.Println("More than two")
  }
}

func NoConditionSwitch() {
  switch x:=2; {
  case x > 2: 
    fmt.Println("Greater than 2")
  case x < 2:
    fmt.Println("Less than 2")
  default: 
    fmt.Println("Equal to 2")
  }
}
