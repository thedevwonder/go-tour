package maps

import "fmt"

func MutateMap() {
  m := make(map[string]int)
  m["Answer"] = 42

  fmt.Println("map value inserted", m["Answer"])

  m["Answer"] = 48 
  fmt.Println("map value updated", m["Answer"])

  delete(m, "Answer")
  fmt.Println("key deleted", m["Answer"])
  
  v, ok := m["Answer"]
  fmt.Println("The Value: ", v, "Present?", ok)
}
