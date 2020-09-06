package main

import "fmt"

func main() {
  fmt.Println("Entry point: ")
  switch_stmt()
}

func count_loop() {
  // for init; condition; post
  for i := 0; i <= 25; i++ {
    // initialises i if it is less thna 25 adn then adds 1 to i until it = 25
    if (i % 5 == 0) {
      fmt.Println(i)
      // if i is a multiple of 5 print it
      }
    }

    fmt.Println("\n")

    // nested loop
    for j := 0; j <= 5; j++ {
      // outer loop
      for n := 0; n <= 5; n++ {
        // inner loop
        fmt.Println("outer loop:", j, "inner loop:", n)
      }
      fmt.Println("\n")
      // every time it does one set a new line
  }
}


func condition_loop() {
  for 1 == 1 {
    // infinite loop that can only be broken
    fmt.Println("hello every body")
    break
  }
  var x int
  // initialise variable outside of the loop
  for x <= 10 {
    fmt.Printf("%d", x)
    x++
    // does the adding within the loop
  }
  fmt.Println("")
  // to make cli look neat

  for i := 0; i <= 122; i++ {
    // loop that goes through ascii identifiers then uses that to print the actual
    // chars
    fmt.Printf("%v \t %#U \t \n", i, i)
  }

  for k := 0; k <= 10; k++ {
    if (k % 2 == 0) {
      fmt.Println(k)
    }
    continue
    // if it is not even restart loop
  }
}


func ifelse() {
  num := 32
  if (num == 32) {
    // need brackets for if not for
    fmt.Println("the value of num is 32")

  } else {

    // else has got to be indented
    fmt.Println("the value of num isn't 32")
  }

}

func switch_stmt() {
  var str string = "5362"
  // initialises a string to be in used in switch statement

  switch {
    // no staement afterwards refers to anything that evaluates to true
  case (str == "I'm a string"):
    fmt.Println("str is a string")
    // string has cases i.e an if statement

  case (str == "I'm not a string"):
    fmt.Println("str says it is not a string")

  default:
    fmt.Println("str is", str)
    // default case is a fall through if none of the cases are true
  }

  switch "hi" {
    // switch using a value evaluates with the value not whether it is true
  case "hello", "hi":
    fmt.Println("Hello, how are you?")
    // if the value in case is equal to value at switch declaration
  case "goodbye":
    fmt.Println("See you later!")

  default:
    fmt.Println("Greetings!")
  }
}
