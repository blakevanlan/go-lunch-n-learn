package main 

import (
   "fmt"
   "math"
)

func pow(x, n, lim float64) float64 {
   if v := math.Pow(x, n); v < lim {
      return v
   } else {
      fmt.Printf("%g >= %g\n", v, lim)
   }
   // no longer have access to v
   return lim
}

func main() {
   fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
   sum := 0

   fmt.Println()
   for i := 0; i < 10; i++ {
      sum += i
   }
   fmt.Println(sum)
}