package basics

import "fmt"

func split (sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func NamedReturn () {
  fmt.Println(split(17))
}
