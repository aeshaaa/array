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

	// create a array , 5 assign values of type int , each index postion & also range function use
	arr1 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr1)

	for i, v := range arr1 { // use range function
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", arr1) // define Type
	

}
