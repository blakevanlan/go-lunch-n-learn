package main

import "fmt"

func main() {
   c := make(chan int, 2) // 2 is the length of the buffer
   c <- 1
   c <- 2
   c <- 3
   fmt.Println(<-c)
   fmt.Println(<-c)
}