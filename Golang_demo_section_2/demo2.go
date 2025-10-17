package main

import "fmt"

func main() {
	const empAge int = 35
	empAge = 36
	fmt.Println(empAge)
}

// In Go, constants are immutable, meaning their values cannot be changed once set. Attempting to reassign a value to a constant will result in a compilation error.
