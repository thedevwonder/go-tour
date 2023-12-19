package basics

import "fmt"

func CallDefer() {
  defer fmt.Println("World")
  fmt.Println("Hello")
}

func StackDefer() {
  for i := 0; i < 10; i++{
    defer fmt.Println(i)
  }
}
