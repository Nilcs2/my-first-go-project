package main

import "fmt"

func main() {
	var emp_name = [4]string{"Nilesh", "Pratishruti", "Bruce", "Bella"}
	var emp_age = [4]int{35, 35, 10, 9}
	fmt.Println(emp_name, emp_age)
	fmt.Printf("Type of emp_age is: %T, Type of emp_name is: %T\n", emp_age, emp_name)
	for i, name := range emp_name {
		fmt.Printf("Age of %s is %d\n", name, emp_age[i])
	}
	fmt.Printf("Age of %s is %d\n", emp_name[0], emp_age[0])
}
