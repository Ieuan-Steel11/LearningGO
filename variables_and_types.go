package main

import "fmt"


func main() {
  fmt.Println("Entry Point: \n")
}

func short_declaration_operator() {
  x := 34
  fmt.Println(x)
  // short declaration requires no type or var keyword

  x = 99080
  fmt.Println(x)
  // allows you to re-assign

  y := 196 + 6789
  fmt.Printf("format print: %d \n", y)
  // stores the evaluation of the statement
}


var global_variable int = 65
// use full declaration for global varibales

func full_declaration() {
  var i int = 56
  fmt.Println(i)
  // full declaration involves using var then name to initialise
  // then a value CAN BE assigned
  // use var keyword for GLOBAL VARIABLES

  var s string
  fmt.Println("string: ", s)
  // automatically asigned false value for boolean but would be 0  for int value
  // would be an empty string also
  // var only declared

  s = "this string now has a value"
  fmt.Println("string: ", s)
  // value now assigned to the variable and is no longer a false or empty value
                                  // GO is statically typed
  // s = 56
  // fmt.Println(s)

  // variable s throws type error because it is a string and has been given an
  // integer value having been declared as a string
}

func types() {
  fmt.Printf("type of variable 'global_variable': %T \n\n", global_variable)
  // %T is for the type of a varibale
  
  var integer int = 5
  // %d
  var decimal float32 = 1.11
  // %.{decimal places}f
  var boolean bool = true
  // %t
  var a_string string = "I am a string"
  // %s

  // main types in go

  fmt.Printf("Types: \n integer: %d \n float: %.2f \n boolean: %t \n string: %s \n",
            integer, decimal, boolean, a_string)

}
