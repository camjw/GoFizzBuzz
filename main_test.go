package main

import "testing"
import "strconv"

func TestMainDivisibleByFifteen(t *testing.T) {
  got := run(30, nil)
  want := "FizzBuzz"
  if got != want {
    t.Errorf("got %.s want %.s", got, want)
  }
}

func TestMainDivisibleByFive(t *testing.T) {
  got := run(25, nil)
  want := "Buzz"
  if got != want {
    t.Errorf("got %.s want %.s", got, want)
  }
}

func TestMainDivisibleByThree(t *testing.T) {
  got := run(81, nil)
  want := "Fizz"
  if got != want {
    t.Errorf("got %.s want %.s", got, want)
  }
}

func TestMainDivisibleByNone(t *testing.T) {
  got := run(101, nil)
  want := "101"
  if got != want {
    t.Errorf("got %.s want %.s", got, want)
  }
}

func TestMainErrorHandling(t *testing.T) {
  input, err := strconv.Atoi("Fifteen")
  got := run(input, err)
  want := "Error, only pass integer as argument"
  if got != want {
    t.Errorf("got %.s want %.s", got, want)
  }
}
