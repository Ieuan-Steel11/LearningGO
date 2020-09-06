package main

import "fmt"

func main() {
  fmt.Println("Entry Point: ")
}

func ex1() {
  fmt.Println(" Excercise 1: ")

  x, y, z := 42, "james bond", true
  // define the vars to be printed

  fmt.Printf("  First Method: \n   %d \n   %s \n   %t \n", x, y, z)
  // print using one line

  fmt.Println("  Second Method: ")
  // print using multiple lines
  fmt.Println("  ", x)
  fmt.Println("  ", y)
  fmt.Println("  " , z)
}

func ex2() {
  var x int
  var y string
  var z bool
  // declare the variables using var

  fmt.Printf("Zero Values: \n Int: '%d' \n String: '%s' \n Bool: '%t' \n", x, y, z)
  // print their 0 values

}


func ex3() {
  x, y, z := 42, "james bond", true
  // define the variables
  str := fmt.Sprintf("%d, %s, %t", x, y, z)
  // puts them all in a single formatted string
  fmt.Println(str)
}


func ex4() {
  type number int
  var x number
  // creates new type number with an underlying type of int and declares var of that type
  fmt.Printf("Value: %d \nType: %T", x ,x)
  // print out 0 value of var
  x = 42
  // assign value
  fmt.Printf("\nAssigned Value: %d\n", x)
}


func ex5() {
  type number int
  var x number

  fmt.Printf("Value: %d \nType: %T", x ,x)
  // print out 0 value of var
  x = 42
  // assign value
  fmt.Printf("\nAssigned Value: %d\n \n", x)

  y := int(x)
  // converts to unerlying type
  // assigns to new var

  fmt.Printf("Converted var Value: %d \nConverted var Type: %T\n", y ,y)
}
