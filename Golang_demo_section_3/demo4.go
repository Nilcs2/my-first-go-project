package main

import "fmt"

func main() {
	var heroDetails = map[string]string{"Name": "Bruce", "Age": "10", "City": "Gotham", "Designation": "Hero"}
	fmt.Println(heroDetails)
	var dogDetails = make(map[string]string)
	dogDetails["Name"] = "Bella"
	dogDetails["Age"] = "9"
	dogDetails["City"] = "Metropolis"
	dogDetails["Designation"] = "Superhero Pet"
	fmt.Println(dogDetails)
	fmt.Printf("Type of heroDetails is: %T, Type of dogDetails is: %T\n", heroDetails, dogDetails)
	fmt.Println("Hero Name is:", heroDetails["Name"])
}
