package main

import (
  "fmt"
)

func intersection(a, b []int) []int {
  counter := make(map[int]int)
  var result []int

  for _, elem := range a {
    if _, ok := counter[elem]; !ok {
      counter[elem] = 1
    } else {
      counter[elem] += 1
    }
  }
  for _, elem := range b {
    if count, ok := counter[elem]; ok && count > 0 {
      counter[elem] -= 1
      result = append(result, elem)
    }
  }
  return result
}

func main() {

  a := []int{37, 5, 1, 2}
  b := []int{6, 2, 4, 37}
  // [2, 37]
  fmt.Printf("%v\n", intersection(a, b))
  a = []int{1, 1, 1}
  b = []int{1, 1, 1, 1}
  // [1, 1, 1]
  fmt.Printf("%v\n", intersection(a, b))
}