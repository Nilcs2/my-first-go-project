package main

import "fmt"

func main() {
	var name string = "Pratishruti"
	fmt.Printf("value: %v, type: %T\n", name, name)
	var myfloat float32 = 3.14
	fmt.Printf("value: %v, type: %T\n", myfloat, myfloat)
	// signed integer
	var signedInt int = 35
	fmt.Printf("Signed integer (int): Value = %v, Type = %T\n", signedInt, signedInt)
	// boolean type
	var isActive bool = true
	fmt.Printf("Boolean: Value = %v, Type = %T\n", isActive, isActive)
}
