package main

import "fmt"

type Vertex struct {
   Lat, Long float64
}

var m = map[string]Vertex{
   "Google": Vertex{
      37.42202, -122.08408,
   },
}

func main() {
   fmt.Println(m)

   // Add
   m["Softek"] = Vertex{ 38.967734, -94.63784}
   fmt.Println(m)

   // Delete
   delete(m, "Softek")
   fmt.Println(m)

   // Check for existance
   _, ok := m["Softek"]
   fmt.Println("Softek Present?", ok)
}
