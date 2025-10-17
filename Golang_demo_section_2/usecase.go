package main

import "fmt"

func main() {
	var username, password string
	fmt.Println("User Registration")
	fmt.Println("Enter Username: ")
	fmt.Scanln(&username)
	fmt.Println("Enter Password: ")
	fmt.Scanln(&password)
	fmt.Println("Username is :", username)
	fmt.Println("The password is :" + "abcd" + password + "wxyz")

}
