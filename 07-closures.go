package main

import "fmt"

func fibonacci() func() int {
   pSum, sum := 0, 1

   return func() int {
      pSum, sum = sum, sum + pSum
      return sum
   }
}

func main() {
   f := fibonacci()

   fmt.Println(0)
   fmt.Println(1)

   for i := 0; i < 20; i++ {
      fmt.Println(f())
   }
}