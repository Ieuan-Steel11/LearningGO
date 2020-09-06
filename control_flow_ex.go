package main

import "fmt"

func main() {
  fmt.Println("Entry point: ")
  ex9()
}

func ex1() {
  // print every num form 1 - 10000
  for i := 0; i <= 10000; i++ {
    fmt.Println(i)
  }
}

func ex2() {
  // print every rune code in the alphabet 3 times
  for i := 0; i <= 122; i++ {
    for j:= 0; j < 3; j++ {
      fmt.Printf("%#U", i)
      // prints each cahracter with that ascii code
    }
    fmt.Println()
    // gives a new line between
  }
}

func ex3() {
  // print years I have been aliv e
  i := 2005
  // stepper var
  for i <= 2020 {
    fmt.Println(i)
    i++
  }
}

func ex4() {
  // print out years I have been alive
  i := 2005
  for {
    if (i > 2020) {
      break
    }
    // as long as it isn  not 2020 yet
    fmt.Println(i)
    i++
    continue
    // print year and add one to year
  }
}

func ex5() {
  for i := 10; i <= 100; i++ {
    // every num between 10 and 100
    fmt.Println(i % 4)
  }
}

func ex6() {
  // use an if statement
  if (1 ==1) {
    fmt.Println(true)
  }
}

func ex7() {
  // use if else if and else
  num := 21
  if (num == 32) {
    fmt.Println("num is 32")
  } else if (num == 21) {
      fmt.Println("num is 21")
  } else {
      fmt.Println("num isn't 32")
  }
}

func ex8() {
  // use a switch
  switch "bye" {
    // case depends on value of var after switch key word
  case "hello", "hi":
    fmt.Println("hello")
  case "bye":
    fmt.Println("goodbye")
  }
}

func ex9() {
  // use a switch
  switch {
  case 1 == 2:
    fmt.Println("one equals two")
  default:
    fmt.Println("one doesn't equal two")
  }
}
