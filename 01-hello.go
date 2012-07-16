package main 

import(
   "fmt"
   "math"
)

const SomeNumber = 42.2
var x, y, z = 1, "two", "three"

func add(x, y float64) float64 {
   return x + y
}

func swap(a, b string) (newA, newB string) {
   newA = b
   newB = a
   return
}

func main() {
   sum := add(SomeNumber, math.Pi)
   newY, newZ := swap(y, z)

   fmt.Println(sum)
   fmt.Println(x, y, z)
   fmt.Println(newY, newZ)
}

/*
   Basic types
   -----------
   bool
   string
   int, int8, int16, int32, int64
   uint, uint8, uint16, uint32, uint64, uintptr
   byte - alias for uint8
   rune - alias for int32
   float32, float 64
   complex64, complex128
*/