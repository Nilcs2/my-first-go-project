package main

import "fmt"

func main() {
	var greetings, company string = "Hello", "Nilesh"
	var year int = 2025
	fmt.Printf("%s %s, welcome to the year %d!\n", greetings, company, year)
	fmt.Println(greetings)
	fmt.Println(company)
}

// In Go, you can declare multiple variables of the same type in a single line by separating them with commas. Here, both 'greetings' and 'company' are declared as strings and assigned values "Hello" and "Nilesh" respectively.
