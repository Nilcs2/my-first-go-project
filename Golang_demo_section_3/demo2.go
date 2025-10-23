package main

import "fmt"

func main() {
	numArr := [5]int{1, 2, 3, 4, 5}
	var s []int = numArr[1:4]
	fmt.Println("The original array:", numArr)
	fmt.Println("The sliced array:", s)
	fmt.Println("Length of sliced array is :", len(s))
	fmt.Println("Capacity of the slice array is :", cap(s))

	slice := []string{"Golang", "Nilesh", "Python", "Pratishruti"}
	fmt.Println("The length of the slice is :", len(slice))
	fmt.Println("The capacity of the slice is :", cap(slice))
}
