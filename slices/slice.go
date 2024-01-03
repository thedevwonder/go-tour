package slices

import "fmt"

func CreateSlice() {
  arr := [6]int{1,2,3,4,5,6}
  var arr_slice []int = arr[1:4]
  fmt.Println(arr_slice)
}

func SliceAsRef() {
  names := [4]string{
    		"John",
		"Paul",
		"George",
		"Ringo",
	   }
  fmt.Println(names)
  a := names[0:2]
  b := names[1:3]
  fmt.Println(a, b)
  b[0] = "XXX"
  fmt.Println(a, b)
  fmt.Println(names)
}
