package main

import "fmt"


func main() {
  fmt.Println("Entry Point: \n")
  iotas()
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

func zero_value() {
  var str string
  var floating float32
  var i int
  var b bool
  // declare variables without a value
  // meaning thye have a 0 zero_value

  fmt.Printf("zero values: \n string: '%s' \n float: %.2f \n int: %d \n bool: %t \n",
              str, floating, i, b)
}

func format_printing() {
  var hello int = 5
  // initialise an int variable

  fmt.Printf("value of integer: %d \nbinary of integer %b \nhexadecimal of integer: %#x",
              hello, hello, hello)
  // diferrent ways of printing the representations of the variable

}

func own_types() {
  type num int
  var a num = 32
  var b int
  // lets you define your own type with an underlying type

  fmt.Println("value:", a)
  fmt.Printf("Type: %T\n", a)
  // prints the actual value and the type which is now num instead of int
  // but has an underlying type of int

  b = int(a)
  // converting between a type with the same underlying type like num and it

  fmt.Println("converted value:", b)
  fmt.Printf("converted Type: %T\n", b)
  // prints teh value and type of the converted variable
}


func constants() {
  const constant = "I AM CONSTANT"
  // you can define constants in go with const key word
  // type is gotten by inference

  const (
    a int8 = 42
    b int8 = 35
    c int8 = 72
    // these are typed constants meaning that you can't mix and match them between
    // custom types whereas unsigned constants can do this
  )
  // define constants in groupes
  fmt.Printf("Constants: \n a: %d \n b: %d \n c: %d \n constant: %s \n",
              a, b, c, constant)
}


func iotas() {
  const (
    a = iota
    b
    c
    // starts from 0 at a and auto increments
    // each time iota keyword is used or const keyword is used it is reset
  )
  fmt.Printf("iotas: \n a: %d \n b: %d \n c: %d \n \n", a, b, c)

  // iota achieves this by bit shifting
  // i.e 1 in bianry is 0000 0001 2 is 0000 0010 3 is 0000 0011
  // it shifts each bit in the byte to increment

  x := 3
  y := x << 2
  // shifts 4 by one and stores in y
  // x is the base of the system i.e 3^0 3^1 3^2

  fmt.Printf("x: %d \t %b \nx bit shifted: %d \t %b \n", x, x, y, y)

  // using iota for kb, mb, gb

  const (
    _ = iota
    // ignores the first variable but iota is there so it knows to shift
    kb = 1 << (iota * 10)
    mb = 1 << (iota * 10)
    gb = 1 << (iota * 10)
    // uses the iota to shift the shifting already happening in the VARIABLES

  )
  fmt.Printf("kb: %d \t\t %b \nmb: %d \t\t %b \ngb: %d \t\t %b \n ",
              kb, kb, mb, mb, gb, gb)
}
