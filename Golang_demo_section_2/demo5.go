package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Enter your name and age: ")
	fmt.Scan(&name, &age)
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}

// In this code, we declare two variables, 'name' of type string and 'age' of type int. We then use fmt.Scan to read user input from the console and store the values in these variables. Finally, we print a greeting message that includes the user's name and age.
