package main

import "fmt"

func main() {
   p := []int{2, 3, 5, 7, 11, 13}
   fmt.Println("p ==", p)

   for i := 0; i < len(p); i++ {
      fmt.Printf("p[%d] == %d\n", i, p[i])
   }

   fmt.Println("p[1:4] ==", p[1:4])
   fmt.Println("p[:3] ==", p[:3]) // implied low index is 0
   fmt.Println("p[4:] ==", p[4:]) // impoled high index is len(s)

   a := make([]int, 5)
   b := make([]int, 0, 5)
   c := b[:3]
   c[0], c[1], c[2] = 1, 2, 3
   d := c[2:5]

   printSlices("a", a)
   printSlices("b", b)
   printSlices("c", c)
   printSlices("d", d)     
}

func printSlices(s string, x []int) {
   fmt.Printf("%s len=%d cap=%d %v\n",
      s, len(x), cap(x), x)
}