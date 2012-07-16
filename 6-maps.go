package main

import "fmt"

type Vertex struct {
   Lat, Long float64
}

var m map[string]Vertex

func main() {
   m = make(map[string]Vertex)
   m["Bell Labs"] = Vertex{
      40.68433, 74.39967,
   }
   n := map[string]Vertex{
      "Google": Vertex{
         37.42202, -122.08408,
      },
   }
  
   fmt.Println(m);
   fmt.Println(n);
   fmt.Println(m["Bell Labs"])
}