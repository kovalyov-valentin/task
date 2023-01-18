package main

import (
      "fmt"
)

// Set Difference: A - B
func Difference(a, b []int) (diff []int) {
      m := make(map[int]bool)

      for _, item := range b {
              m[item] = true
      }

      for _, item := range a {
              if _, ok := m[item]; !ok {
                      diff = append(diff, item)
              }
      }
      return
}

func main() {
      var a = []int{1, 2, 3, 4, 5}
      var b = []int{2, 3, 5, 7, 11}
      fmt.Println(Difference(a, b))
}