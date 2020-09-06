package main

import "fmt"

func main() {
  fmt.Println("Entry point: ")
  ex6()
}

func ex1() {
  x := 578
  // initalise variable to be printed

  fmt.Printf("Denary: %d \t Binary: %b \t Hex: %#x \n", x, x, x)
}

func ex2() {
  a := 56 == 65
  b := 56 <= 65
  c := 56 >= 65
  d := 56 != 65
  e := 56 > 65
  f := 56 < 65
  // use every operator then print results
  fmt.Println(a, b, c, d, e, f)
}

func ex3() {
  const (
    a = iota
    b
    c
    d
  )
  const number int = 65
  // initialise constants to print
  fmt.Println("Untyped:", a, b, c, d)
  fmt.Println("Typed:", number)
}

func ex4() {
  var num int = 1748
  // intialise var
  fmt.Printf("Denary: %d \t Binary: %b \t Hex: %#x \n", num, num, num)

  shifted_num := num << 1
  // shift by one and print
  fmt.Printf("Denary: %d \t Binary: %b \t Hex: %#x \n", shifted_num, shifted_num,
              shifted_num)
}

func ex5() {
  raw_string_literal := `   This is a \t           raw
                                      string
                                literal`
  // raw string literal doesn't account for escape characters it's just literally
  // everything between the marks
  fmt.Println(raw_string_literal)

}

func ex6() {
  const (
    a = 2020 + iota
    b = 2020 + iota
    c = 2020 + iota
    d = 2020 + iota
  )
  fmt.Println(a, b, c, d)
  // create constants for the next 4 years and print
}
