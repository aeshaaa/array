package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Println(x)

	arr := [3]int{3, 5, 2}
	fmt.Println(arr)

	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println(numbers[0:2])
}

// 2 types of array
