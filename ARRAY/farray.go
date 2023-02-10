package main

import "fmt"

func main() {
	var b = []int{1000, 100, 10, 1}
	var avg float32
	avg = getavg[b,5]
	fmt.Printf("avg is %f", avg)
}

func getavg(arr []int, size int) float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum / size)
	return avg
}
