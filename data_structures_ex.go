package main

import "fmt"

func main() {
  fmt.Println("Entry point:")
  ex8_9_10()
}

func ex1() {
  arr := [5]int{1, 2, 3, 4, 5}
  // created array with composite literal

  for index, value := range arr {
    fmt.Println("At index:", index, "value:", value)
    // prints all values in array
  }
  fmt.Printf("Type of arr: %T\n", arr)
}

func ex2() {
  slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
  // slice with 10 values

  for index, value := range slice {
    fmt.Println("At index:", index, "value:", value)
    // prints all values in slice
  }
  fmt.Printf("Type of slice: %T \n", slice)
}

func ex3() {
  slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
  // creating a slice

  fmt.Println(slice[:5])
  fmt.Println(slice[5:])
  fmt.Println(slice[2:7])
  fmt.Println(slice[1:6])
  // printing slices
}

func ex4() {
  slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
  // creating a slice

  slice = append(slice, 52)
  fmt.Println(slice)
  // add number 52

  nums_to_add := []int{53, 54, 55}
  slice = append(slice, nums_to_add...)
  fmt.Println(slice)
  // adds three more numbers

  nums_to_add = []int{56, 57, 58, 59, 60}
  slice = append(slice, nums_to_add...)
  fmt.Println(slice)
  // adding more numbers
}

func ex5() {
  slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

  slice = append(slice[0:3], slice[6:]...)
  fmt.Println(slice)
}

func ex6() {
  counties := []string{"Devon", "Cornwall", "Herefordshire", "Pembrokeshire", "Essex", "Wessex"}
  i := 0
  for i < len(counties) {
    fmt.Println("Index:", i, "County", counties[i])
    i++
  }
  fmt.Println("Amount of counties", i)
}

func ex7() {
  slice1 := []string{"james", "bond", "shaken not stirred"}
  slice2 := []string{"miss", "moneypenny", "hello james"}
  slices := [][]string{slice1, slice2}

  for index, value := range slices {
    fmt.Println(index, value)
    for i, v := range value {
      fmt.Println(i, v)
    }
  }
}

func ex8_9_10() {
  m := map[string][]string{
    "bond_james": {"shaken, not stirred", "Martinis"},
    "no_dr": {"being evil", "ice cream",  "sunsets"},
  }
  for key, value := range m {
    fmt.Println(key, value)
    for i, v := range value {
      fmt.Println(i, v)
    }
  }
  m["me"] = []string{"annoyed", "about", "this"}
  delete(m, "bond_james")

  for k, v := range m {
    fmt.Println(k, v)
  }
}
