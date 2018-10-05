package main

import fb "github.com/camjw/fizzbuzz/fizzbuzz"
import "fmt"
import "os"
import "strconv"

func run(signal int, err error) (string) {
  if err != nil {
     return("Error, only pass integer as argument")
  }
  var output = fb.FizzBuzz{
    Signal: signal,
  }
  return(output.Output())
}

func main() {
  var input int
  input, err := strconv.Atoi(os.Args[1])
  fmt.Println(run(input, err))
}
