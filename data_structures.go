package main

import (
  "time"
  "fmt"
)

func main() {
  fmt.Println("Entry point: ")
  slices()
}


func arrays() {
  // for arrays you have to assign a datatype and a size when declaring
  var arr [5]int
  // size in square brackets type afterwards
  fmt.Println(arr)
  // zero indexed
  arr[0] = 234
  // assign a value to item at index 0
  fmt.Println(arr)
  fmt.Println(len(arr))
  // there is also a len function in go

  // in go slices are used instead of arrays unless you have a detailed
  // layout to implement thta doesn't change
}

func slices() {
  // x := type{values}
  slice := []int{4, 2, 5, 1, 6, 3, 789 ,7809, 456, 687, 6789, 324, 2134, 6543, 456, 543, 4327564, 5432, 43245321, 432543245626, 213, 54632, 765, 234, 9876, 234432, 87648765, 423, 7653767, 7653, 7635, 7653, 987, 5463, 231, 67543, 765, 98, 456, 223, 543, 6534, 6543, 654, 654, 654, 87776, 87646, 56767, 654546, 3454, 543, 534543, 543543354, 53454334345354, 343, 3465}
  // based of an array it doesn't take a given size but values at th end
  fmt.Println("Original array:", slice, "\n")
  // testing a bubble sort
  swap_count := 0
  loop_count := 0
  // var to count amount of sorts required
  start := time.Now()
  // start measuring time
  for j := 0; j < (len(slice) - 1); j++ {
    for i := 0; i < (len(slice) - 1); i++ {
      // inner loop does one run through
      if (slice[i] > slice[i + 1]) {
        // if the one before is greater than the one after
        slice[i], slice[i + 1] = slice[i + 1], slice[i]
        swap_count++
        loop_count++
        fmt.Printf("after %d swap(s) and %d iteration(s) %d\n \n", swap_count, loop_count, slice)
        continue
      }
      loop_count++
      continue
      // otherwise move on to next bubble
    }
  }
  duration := time.Since(start)
  // stop measuring time
  fmt.Printf("An array with length %d \n %d swaps \n %d iterations \n %s to execute\n", len(slice), swap_count, loop_count, duration)
  fmt.Println("sorted array: \n", slice, "\n")

  // iterating with range clause
  for index, value := range slice {
    fmt.Println(index, value)
  }
  // appending to a slice
  // append(slice, elements to add/slice to add...)
  // appending another slice requires ... operator after 2nd slice which unpacks
  // the slice into integers which is what the function needs as arguements
  slice = append(slice, 7, 8, 9)
  fmt.Println(slice)

  nums := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
  // appending another slice
  slice = append(slice, nums...)
  fmt.Println(slice)

  // deleting from a slice using append
  nums = append(nums[:2], nums[5:]...)
  // gets rid of the indexes that aren't mentioned in append i.e 3, 4, 5
  fmt.Println(nums)
}

func multi_dimensional_slice() {
  name := []string{"Ieuan", "Steel", "15"}
  name2 := []string{"James", "Steel", "13"}
  // define slices to go into 2d slice
  nms := [][]string{name, name2}
  // initialises 2d string
  fmt.Println(nms)
}

func maps() {
  // map := map[key_type]value_type{entries}
  // like a python dictionary
  m := map[string]int{
    "Ieuan": 15,
    "James": 13,
    "Elaine": 48,
  }
  // if there isn't the entered key in the dictionary returns 0
  fmt.Println(m["Ieuan"])

  m["Stuart"] = 50
  // adding to a map

  if v, ok := m["Stuart"]; ok {
    // if there is the key stuart in the map
    fmt.Println(" Value:", v, "\n", "exists:",ok)
  } else {
    fmt.Println("Stuart isn't in the map")
  }

  for key, value := range m {
    // iterate iver each key and value in a map
    fmt.Println("Key:", key, "Value:", value)
  }

  delete(m, "Stuart")
  // can delete entries that don't exist

  fmt.Println(m)
}
