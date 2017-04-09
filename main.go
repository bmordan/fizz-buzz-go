package main

import "fmt"

func main() {
  for i := 1; i < 101; i++ {
    var result string
    if i % 3 == 0 && i % 5 == 0 {
      result = "FizzBuzz"
    } else if i % 5 == 0 {
      result = "Buzz"
    } else if i % 3 == 0 {
      result = "Fizz"
    }
    if len(result) == 0 {
      fmt.Println(i)
    } else {
      fmt.Println(result)
    }
  }
}
