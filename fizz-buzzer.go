package fizzBuzzer

func (i) {
  var result string
  if i % 3 == 0 && i % 5 == 0 {
    result = "FizzBuzz"
  } else if i % 5 == 0 {
    result = "Buzz"
  } else if i % 3 == 0 {
    result = "Fizz"
  }
  if len(result) == 0 {
    return i
  } else {
    return result
  }
}
