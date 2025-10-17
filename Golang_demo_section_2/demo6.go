package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Enter your name and age: ")
	count, err := fmt.Scanf("%s, %d", &name, &age)
	fmt.Println("Count", count)
	fmt.Println("Error", err)
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}
