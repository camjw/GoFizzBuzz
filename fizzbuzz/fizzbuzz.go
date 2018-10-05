package fizzbuzz

import "strconv"

type FizzBuzz struct {
  Signal int
}


func (f FizzBuzz) Output() (string) {
  if f.Signal%15 == 0 {
    return "FizzBuzz"
  } else if f.Signal%5 == 0 {
    return "Buzz"
  } else if f.Signal%3 ==0 {
    return "Fizz"
  } else {
    return strconv.Itoa(f.Signal)
  }
}
