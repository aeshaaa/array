package main

import (
	"fmt"
)

func main() {
	arr1 := [2]int{2, 3}
	arr2 := [...]int{2, 3}
	fmt.Println(arr1 == arr2)
}

// output of the below code ...(T ya F)
