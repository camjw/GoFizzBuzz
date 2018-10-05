package main

import fb "github.com/camjw/fizzbuzz/fizzbuzz"
import "fmt"
import "os"
import "strconv"

func main() {
  var input int
  input, err := strconv.Atoi(os.Args[1])
  if err != nil {
     fmt.Println("Error, only pass integer as argument")
  }
  var output = fb.FizzBuzz{
    Signal: input,
  }
  fmt.Println(output.Output())
}
