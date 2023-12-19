package basics

import "fmt"

func Loop() {
  sum := 0

  for i := 0; i < 10; i++ {
    sum += i
  }

  fmt.Println(sum)
}

func Loop2() {
  sum := 1

  for ; sum < 1000; {
    sum += sum
  }

  fmt.Println(sum)

}

func Loop3() {
  sum := 1
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)
}

func InfiniteLoop() {
  for {
  }
}
