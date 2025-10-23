package main

import "fmt"

func main() {
	numSlice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(numSlice1[3])
	numSlice1[2] = 99
	fmt.Println(numSlice1)
	numSlice2 := []int{6, 7, 8, 9, 10}
	slice1 := append(numSlice1, numSlice2...)
	fmt.Println("After appending:", slice1)
}
