package main 

import (
	"fmt"
)

func main() {
	arrays := []int{1, 2, 3, 4, 5}
	fmt.Println(unique(arrays))
}

func unique(nums []int) bool {
	unique := make(map[int]bool)
	for _, num := range nums {
		_, ok := unique[num]
		if ok {
			return false 
		}
		unique[num] = true
	}
	return true
}