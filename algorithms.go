package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("Entry point: ")
  slice := []int{123, 312, 1342, 4, 1, 312, 413, 432, 5364532, 654, 3145, 5413, 5431,
    567543, 765, 76531, 54, 354, 6542, 44312, 44312, 12345, 545341, 5431, 5431,
    5423, 7654, 7654, 9875, 143, 78435, 9875, 2354, 1543265, 543231245, 543254,
    53425432, 5432, 5432, 6543, 4321, 65243, 754, 14536542, 6542, 6542, 6254,
    6542, 6542, 6542, 6542, 6542, 6542, 6542, 1453, 7653, 95876, 9857, 4135,
    54312, 56473, 658765, 89765, 876, 2345, 564, 543, 5123, 3145, 5314543, 5431}
  // based of an array it doesn't take a given size but values at th end
  slice = append(slice, make([]int, 2500)...)
  // create slice to be sorted
  bubblesort(slice)
}

func bubblesort(slice []int) {
  // x := type{values}
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
}

func experiment() {
  swap_count := 0
  loop_count := 0
  // var to count amount of sorts required
  start := time.Now()
  // start measuring time
  for j := 0; j < (len(slice) - 1); j++ {
  
  }
  duration := time.Since(start)
  // stop measuring time
  fmt.Printf("An array with length %d \n %d swaps \n %d iterations \n %s to execute\n", len(slice), swap_count, loop_count, duration)
}
