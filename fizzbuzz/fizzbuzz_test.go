package fizzbuzz

import "testing"

func TestDivisibleByFifteen(t *testing.T) {
  fizzbuzz := FizzBuzz{Signal: 30}
  got := fizzbuzz.Output()
  want := "FizzBuzz"
  if got != want {
    t.Errorf("got %.s want %.s", got, want)
  }
}

func TestDivisibleByFive(t *testing.T) {
  fizzbuzz := FizzBuzz{Signal: 110}
  got := fizzbuzz.Output()
  want := "Buzz"
  if got != want {
    t.Errorf("got %.s want %.s", got, want)
  }
}

func TestDivisibleByThree(t *testing.T) {
  fizzbuzz := FizzBuzz{Signal: 273}
  got := fizzbuzz.Output()
  want := "Fizz"
  if got != want {
    t.Errorf("got %.s want %.s", got, want)
  }
}

func TestDivisibleByNone(t *testing.T) {
  fizzbuzz := FizzBuzz{Signal: 103}
  got := fizzbuzz.Output()
  want := "103"
  if got != want {
    t.Errorf("got %.s want %.s", got, want)
  }
}
