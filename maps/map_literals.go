package maps

import "fmt"

func UseMapLiteral() {
  m := map[string]Vertex {
  	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
    }
  fmt.Println(m)
}

func UseMapLiteral2() {
  m := map[string]Vertex {
        "Bell Labs": {
                40.68433, -74.39967,
        },
        "Google": {
                37.42202, -122.08408,
        },
    }
  fmt.Println(m)
}
