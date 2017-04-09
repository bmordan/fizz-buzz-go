package main

import "fmt"
import "github.com/bmordan/fizz-buzz/fizzbuzz"

func main() {
  for i := 1; i < 101; i++ {
    fmt.Println(fizzbuzz.Calculate(i))
  }
}
