package main

import "fmt"

type heroDetails struct {
	Name        string
	Age         int
	City        string
	Designation string
}

func main() {
	var hero1 = heroDetails{Name: "Bruce", Age: 10, City: "Gotham", Designation: "The_Dark_Knight"}
	var hero2 = heroDetails{Name: "Bella", Age: 9, City: "Metropolis", Designation: "Cat_Woman"}
	fmt.Println(hero1)
	fmt.Println(hero2)
	fmt.Printf("Type of hero1 is: %T, Type of hero2 is: %T\n", hero1, hero2)
	fmt.Println("Hero1 Name is:", hero1.Name)
	fmt.Printf("%s's City is: %s\n", hero2.Name, hero2.City)
}
