package main // This tells the Go compiler that this is the main package and it can be executed directly.
import "fmt" // Importing the fmt package to use its functions for formatted Input and Output.

func main() {
	//I am creating a variable for my name here
	var name string = "Nilesh" // The variable 'name' is of type string and is assigned the value "Nilesh".
	// Now I will create the variable age

	age := 35 // The variable 'age' is inferred to be of type int and is assigned the value 35.

	fmt.Println("Hello! my name is", name, "and my age is", age)

}
