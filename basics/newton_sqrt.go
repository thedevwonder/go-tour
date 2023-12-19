package basics

import "fmt"

func Sqrt (x int) (y float64) {
  y = 1.0
  for z:=0; z < 10; z++ {
    y -= (y*y - float64(x)) / (2*y)
  }
  fmt.Printf("sqrt of %v is %v", x, y)
  
  return	
}
