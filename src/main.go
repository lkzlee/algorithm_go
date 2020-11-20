package main

import (
	"algorithm_go/src/dp/_53_maximum_subarray"
	"fmt"
)

func main() {
	fmt.Print("hello\n")
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	max := _53_maximum_subarray.MaxSubArray(nums)

	fmt.Println("max sub array result:", max)
}
