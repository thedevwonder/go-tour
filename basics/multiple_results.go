package basics

import "fmt"

func swap(x, y string) (string, string) {
  return y, x
}

func CallSwap(x, y string) {
  a, b := swap(x, y)
  fmt.Println("x and y are swapped", a, b)
}
